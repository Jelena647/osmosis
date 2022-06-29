package keeper

import (
	"fmt"
	"time"

	db "github.com/tendermint/tm-db"

	"github.com/osmosis-labs/osmosis/v7/x/incentives/types"
	lockuptypes "github.com/osmosis-labs/osmosis/v7/x/lockup/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Returns coins that have been distributed already from the provided gauges
func (k Keeper) getDistributedCoinsFromGauges(gauges []types.Gauge) sdk.Coins {
	coins := sdk.Coins{}
	for _, gauge := range gauges {
		coins = coins.Add(gauge.DistributedCoins...)
	}
	return coins
}

// Returns coins that have not been distributed yet from the provided gauges
func (k Keeper) getToDistributeCoinsFromGauges(gauges []types.Gauge) sdk.Coins {
	// TODO: Consider optimizing this in the future to only require one iteration over all gauges.
	// TODO: Lets make issue for this, requires changing getGaugesFromIterator logic in gauge.go
	coins := k.getCoinsFromGauges(gauges)
	distributed := k.getDistributedCoinsFromGauges(gauges)
	return coins.Sub(distributed)
}

// Utilizes db.iterator to return a list of gauges.
// From these gauges, coins that have not yet been distributed are returned
func (k Keeper) getToDistributeCoinsFromIterator(ctx sdk.Context, iterator db.Iterator) sdk.Coins {
	return k.getToDistributeCoinsFromGauges(k.getGaugesFromIterator(ctx, iterator))
}

// Utilizes db.iterator to return a list of gauges.
// From these gauges, coins that have already been distributed are returned
func (k Keeper) getDistributedCoinsFromIterator(ctx sdk.Context, iterator db.Iterator) sdk.Coins {
	return k.getDistributedCoinsFromGauges(k.getGaugesFromIterator(ctx, iterator))
}

// Moves a gauge that has reached it's start time from an upcoming to an active status.
func (k Keeper) moveUpcomingGaugeToActiveGauge(ctx sdk.Context, gauge types.Gauge) error {
	// validation for current time and distribution start time
	if ctx.BlockTime().Before(gauge.StartTime) {
		return fmt.Errorf("gauge is not able to start distribution yet: %s >= %s", ctx.BlockTime().String(), gauge.StartTime.String())
	}

	timeKey := getTimeKey(gauge.StartTime)
	if err := k.deleteGaugeRefByKey(ctx, combineKeys(types.KeyPrefixUpcomingGauges, timeKey), gauge.Id); err != nil {
		return err
	}
	if err := k.addGaugeRefByKey(ctx, combineKeys(types.KeyPrefixActiveGauges, timeKey), gauge.Id); err != nil {
		return err
	}
	return nil
}

// Moves a gauge that has completed its distribution from an active to a finished status.
func (k Keeper) moveActiveGaugeToFinishedGauge(ctx sdk.Context, gauge types.Gauge) error {
	timeKey := getTimeKey(gauge.StartTime)
	if err := k.deleteGaugeRefByKey(ctx, combineKeys(types.KeyPrefixActiveGauges, timeKey), gauge.Id); err != nil {
		return err
	}
	if err := k.addGaugeRefByKey(ctx, combineKeys(types.KeyPrefixFinishedGauges, timeKey), gauge.Id); err != nil {
		return err
	}
	if err := k.deleteGaugeIDForDenom(ctx, gauge.Id, gauge.DistributeTo.Denom); err != nil {
		return err
	}
	k.hooks.AfterFinishDistribution(ctx, gauge.Id)
	return nil
}

// Returns locks that match the provided lockuptypes QueryCondition AND have yet to be distributed to yet.
// TODO: Why is this exported?
func (k Keeper) GetLocksToDistribution(ctx sdk.Context, distrTo lockuptypes.QueryCondition) []lockuptypes.PeriodLock {
	switch distrTo.LockQueryType {
	case lockuptypes.ByDuration:
		return k.lk.GetLocksLongerThanDurationDenom(ctx, distrTo.Denom, distrTo.Duration)
	case lockuptypes.ByTime:
		// TODO: Shouldn't this panic since we don't support lockuptypes byTime?
		return k.lk.GetLocksPastTimeDenom(ctx, distrTo.Denom, distrTo.Timestamp)
	default:
	}
	return []lockuptypes.PeriodLock{}
}

// Returns locks that match the provided lockuptypes QueryCondition,
// are greater than the provided minDuration, AND have yet to be distributed to.
func (k Keeper) getLocksToDistributionWithMaxDuration(ctx sdk.Context, distrTo lockuptypes.QueryCondition, minDuration time.Duration) []lockuptypes.PeriodLock {
	switch distrTo.LockQueryType {
	case lockuptypes.ByDuration:
		// TODO: Create issue for this
		// Confusingly, synthetic lockups aren't indexed by denom as expected.
		// Thus you have to do this as a hack
		denom := lockuptypes.NativeDenom(distrTo.Denom)
		if distrTo.Duration > minDuration {
			return k.lk.GetLocksLongerThanDurationDenom(ctx, denom, minDuration)
		}
		return k.lk.GetLocksLongerThanDurationDenom(ctx, distrTo.Denom, distrTo.Duration)
	case lockuptypes.ByTime:
		panic("Gauge by time is present, however is no longer supported. This should have been blocked in ValidateBasic")
	default:
	}
	return []lockuptypes.PeriodLock{}
}

// Estimates distribution amount of coins from gauge
// Expectation: gauge is a valid gauge
// filteredLocks are all locks that are valid for gauge
// It also applies an update for the gauge, handling the sending of the rewards.
// (Note this update is in-memory, it does not change state.)
func (k Keeper) FilteredLocksDistributionEst(ctx sdk.Context, gauge types.Gauge, filteredLocks []lockuptypes.PeriodLock) (types.Gauge, sdk.Coins, error) {
	TotalAmtLocked := k.lk.GetPeriodLocksAccumulation(ctx, gauge.DistributeTo)
	if TotalAmtLocked.IsZero() {
		return types.Gauge{}, nil, nil
	}

	remainCoins := gauge.Coins.Sub(gauge.DistributedCoins)
	// remainEpochs is the number of remaining epochs that the gauge will pay out its rewards
	// For a perpetual gauge, it will pay out everything in the next epoch, and we don't make
	// an assumption of the rate at which it will get refilled at.
	remainEpochs := uint64(1)
	if !gauge.IsPerpetual {
		remainEpochs = gauge.NumEpochsPaidOver - gauge.FilledEpochs
	}
	// TODO: Should this return err
	// TODO: Is this all that was needed?
	if remainEpochs == 0 {
		return gauge, sdk.Coins{}, fmt.Errorf("There are no more remaining epochs for this gauge.")
	}

	remainCoinsPerEpoch := sdk.Coins{}
	for _, coin := range remainCoins {
		// distribution amount per epoch = gauge_size / (remain_epochs)
		amt := coin.Amount.QuoRaw(int64(remainEpochs))
		remainCoinsPerEpoch = remainCoinsPerEpoch.Add(sdk.NewCoin(coin.Denom, amt))
	}

	// Now we compute the filtered coins
	filteredDistrCoins := sdk.Coins{}
	if len(filteredLocks) == 0 {
		// If were doing no filtering, we want to calculate the total amount to distributed in
		// the next epoch.
		// distribution in next epoch = gauge_size  / (remain_epochs)
		filteredDistrCoins = remainCoinsPerEpoch
	}
	for _, lock := range filteredLocks {
		denomLockAmt := lock.Coins.AmountOf(gauge.DistributeTo.Denom)

		for _, coin := range remainCoinsPerEpoch {
			// distribution amount = gauge_size * denom_lock_amount / (total_denom_lock_amount * remain_epochs)
			// distribution amount = gauge_size_per_epoch * denom_lock_amount / total_denom_lock_amount
			amt := coin.Amount.Mul(denomLockAmt).Quo(TotalAmtLocked)
			filteredDistrCoins = filteredDistrCoins.Add(sdk.NewCoin(coin.Denom, amt))
		}
	}

	// increase filled epochs after distribution
	gauge.FilledEpochs += 1
	gauge.DistributedCoins = gauge.DistributedCoins.Add(remainCoinsPerEpoch...)

	return gauge, filteredDistrCoins, nil
}

// Stores all of the information for pent up sends for rewards distributions.
// This enables us to lower the number of events and calls to back.
type distributionInfo struct {
	nextID            int
	lockOwnerAddrToID map[string]int
	idToBech32Addr    []string
	idToDecodedAddr   []sdk.AccAddress
	idToDistrCoins    []sdk.Coins
}

// Creates a new distributionInfo struct
func newDistributionInfo() distributionInfo {
	return distributionInfo{
		nextID:            0,
		lockOwnerAddrToID: make(map[string]int),
		idToBech32Addr:    []string{},
		idToDecodedAddr:   []sdk.AccAddress{},
		idToDistrCoins:    []sdk.Coins{},
	}
}

// TODO: Slightly confused here, how does this determine which lockID to add rewards to? Cant a single owner own multiple locks?
func (d *distributionInfo) addLockRewards(owner string, rewards sdk.Coins) error {
	if id, ok := d.lockOwnerAddrToID[owner]; ok {
		oldDistrCoins := d.idToDistrCoins[id]
		d.idToDistrCoins[id] = rewards.Add(oldDistrCoins...)
	} else {
		id := d.nextID
		d.nextID += 1
		d.lockOwnerAddrToID[owner] = id
		decodedOwnerAddr, err := sdk.AccAddressFromBech32(owner)
		if err != nil {
			return err
		}
		d.idToBech32Addr = append(d.idToBech32Addr, owner)
		d.idToDecodedAddr = append(d.idToDecodedAddr, decodedOwnerAddr)
		d.idToDistrCoins = append(d.idToDistrCoins, rewards)
	}
	return nil
}

// Utilizes provided distributionInfo to send coins from the module account to various recipients
func (k Keeper) doDistributionSends(ctx sdk.Context, distrs *distributionInfo) error {
	numIDs := len(distrs.idToDecodedAddr)
	ctx.Logger().Debug(fmt.Sprintf("Beginning distribution to %d users", numIDs))
	err := k.bk.SendCoinsFromModuleToManyAccounts(
		ctx,
		types.ModuleName,
		distrs.idToDecodedAddr,
		distrs.idToDistrCoins)
	if err != nil {
		return err
	}
	ctx.Logger().Debug("Finished sending, now creating liquidity add events")
	for id := 0; id < numIDs; id++ {
		ctx.EventManager().EmitEvents(sdk.Events{
			sdk.NewEvent(
				types.TypeEvtDistribution,
				sdk.NewAttribute(types.AttributeReceiver, distrs.idToBech32Addr[id]),
				sdk.NewAttribute(types.AttributeAmount, distrs.idToDistrCoins[id].String()),
			),
		})
	}
	ctx.Logger().Debug(fmt.Sprintf("Finished Distributing to %d users", numIDs))
	return nil
}

// Runs the distribution logic for a synthetic rewards distribution gauge, and adds the sends to
// the distrInfo struct. It also updates the gauge for the distribution.
// locks is expected to be the correct set of lock recipients for this gauge.
// TODO: Make this code have way more re-use with distribute internal (post-v7).
// TODO: Create issue for the above
func (k Keeper) distributeSyntheticInternal(
	ctx sdk.Context, gauge types.Gauge, locks []lockuptypes.PeriodLock, distrInfo *distributionInfo,
) (sdk.Coins, error) {
	totalDistrCoins := sdk.NewCoins()
	denom := gauge.DistributeTo.Denom

	qualifiedLocks := make([]lockuptypes.PeriodLock, 0, len(locks))
	for _, lock := range locks {
		// See if this lock has a synthetic lockup. If so, err == nil, and we add to qualifiedLocks
		// otherwise it does not, and we continue.
		_, err := k.lk.GetSyntheticLockup(ctx, lock.ID, denom)
		if err != nil {
			continue
		}
		qualifiedLocks = append(qualifiedLocks, lock)
	}

	// If locksum is zero, there are no synthetic locks so we return nil
	lockSum := lockuptypes.SumLocksByDenom(qualifiedLocks, lockuptypes.NativeDenom(denom))

	if lockSum.IsZero() {
		return nil, nil
	}

	remainCoins := gauge.Coins.Sub(gauge.DistributedCoins)
	// If its a perpetual gauge, we set remaining epochs to 1
	// Otherwise is is a non perpetual gauge and we determine how many epoch payouts are left
	remainEpochs := uint64(1)
	if !gauge.IsPerpetual {
		remainEpochs = gauge.NumEpochsPaidOver - gauge.FilledEpochs
	}

	for _, lock := range qualifiedLocks {
		distrCoins := sdk.Coins{}
		for _, coin := range remainCoins {
			lockedCoin, err := lock.SingleCoin()
			if err != nil {
				k.Logger(ctx).Error(err.Error())
				continue
			}
			// distribution amount = gauge_size * denom_lock_amount / (total_denom_lock_amount * remain_epochs)
			// check if the synthlock is qualified for GetLocksToDistribution
			amt := coin.Amount.Mul(lockedCoin.Amount).Quo(lockSum.Mul(sdk.NewIntFromUint64(remainEpochs)))
			if amt.IsPositive() {
				newlyDistributedCoin := sdk.Coin{Denom: coin.Denom, Amount: amt}
				distrCoins = distrCoins.Add(newlyDistributedCoin)
			}
		}
		distrCoins = distrCoins.Sort()
		if distrCoins.Empty() {
			continue
		}
		// Update the amount for that address
		err := distrInfo.addLockRewards(lock.Owner, distrCoins)
		if err != nil {
			return nil, err
		}

		totalDistrCoins = totalDistrCoins.Add(distrCoins...)
	}

	// increase filled epochs after distribution
	err := k.updateGaugePostDistribute(ctx, gauge, totalDistrCoins)
	return totalDistrCoins, err
}

// Runs the distribution logic for a gauge, and adds the sends to
// the distrInfo struct. It also updates the gauge for the distribution.
// locks is expected to be the correct set of lock recipients for this gauge.
func (k Keeper) distributeInternal(
	ctx sdk.Context, gauge types.Gauge, locks []lockuptypes.PeriodLock, distrInfo *distributionInfo,
) (sdk.Coins, error) {
	totalDistrCoins := sdk.NewCoins()
	denom := gauge.DistributeTo.Denom

	// If locksum is zero, there are no locks so we return nil
	lockSum := lockuptypes.SumLocksByDenom(locks, denom)

	if lockSum.IsZero() {
		return nil, nil
	}

	remainCoins := gauge.Coins.Sub(gauge.DistributedCoins)
	// If its a perpetual gauge, we set remaining epochs to 1
	// Otherwise is is a non perpetual gauge and we determine how many epoch payouts are left
	remainEpochs := uint64(1)
	if !gauge.IsPerpetual {
		remainEpochs = gauge.NumEpochsPaidOver - gauge.FilledEpochs
	}

	for _, lock := range locks {
		distrCoins := sdk.Coins{}
		for _, coin := range remainCoins {
			// distribution amount = gauge_size * denom_lock_amount / (total_denom_lock_amount * remain_epochs)
			denomLockAmt := lock.Coins.AmountOfNoDenomValidation(denom)
			amt := coin.Amount.Mul(denomLockAmt).Quo(lockSum.Mul(sdk.NewInt(int64(remainEpochs))))
			if amt.IsPositive() {
				newlyDistributedCoin := sdk.Coin{Denom: coin.Denom, Amount: amt}
				distrCoins = distrCoins.Add(newlyDistributedCoin)
			}
		}
		distrCoins = distrCoins.Sort()
		if distrCoins.Empty() {
			continue
		}
		// Update the amount for that address
		err := distrInfo.addLockRewards(lock.Owner, distrCoins)
		if err != nil {
			return nil, err
		}

		totalDistrCoins = totalDistrCoins.Add(distrCoins...)
	}

	err := k.updateGaugePostDistribute(ctx, gauge, totalDistrCoins)
	return totalDistrCoins, err
}

// Increments the gauge's filled epochs field.
// Also adds the coins that were just distributed to the gauge's distributed coins field.
func (k Keeper) updateGaugePostDistribute(ctx sdk.Context, gauge types.Gauge, newlyDistributedCoins sdk.Coins) error {
	gauge.FilledEpochs += 1
	gauge.DistributedCoins = gauge.DistributedCoins.Add(newlyDistributedCoins...)
	if err := k.setGauge(ctx, &gauge); err != nil {
		return err
	}
	return nil
}

// Takes a gauge along with cached period locks by denom and returns locks that must be distributed to
func (k Keeper) getDistributeToBaseLocks(ctx sdk.Context, gauge types.Gauge, cache map[string][]lockuptypes.PeriodLock) []lockuptypes.PeriodLock {
	// if gauge is empty, don't get the locks
	if gauge.Coins.Empty() {
		return []lockuptypes.PeriodLock{}
	}
	// TODO: FIXME!!!
	// TODO: Make issue for the above
	// Confusingly, there is no way to get all synthetic lockups. Thus we use a separate method `distributeSyntheticInternal` to separately get lockSum for synthetic lockups.
	// All gauges have a precondition of being ByDuration.
	distributeBaseDenom := lockuptypes.NativeDenom(gauge.DistributeTo.Denom)
	if _, ok := cache[distributeBaseDenom]; !ok {
		cache[distributeBaseDenom] = k.getLocksToDistributionWithMaxDuration(
			ctx, gauge.DistributeTo, time.Millisecond)
	}
	// get this from memory instead of hitting iterators / underlying stores.
	// due to many details of cacheKVStore, iteration will still cause expensive IAVL reads.
	allLocks := cache[distributeBaseDenom]
	return FilterLocksByMinDuration(allLocks, gauge.DistributeTo.Duration)
}

// Distribute coins from an array of gauges to all eligible locks
func (k Keeper) Distribute(ctx sdk.Context, gauges []types.Gauge) (sdk.Coins, error) {
	distrInfo := newDistributionInfo()

	locksByDenomCache := make(map[string][]lockuptypes.PeriodLock)
	totalDistributedCoins := sdk.Coins{}
	for _, gauge := range gauges {
		filteredLocks := k.getDistributeToBaseLocks(ctx, gauge, locksByDenomCache)
		// send based on synthetic lockup coins if it's distributing to synthetic lockups
		var gaugeDistributedCoins sdk.Coins
		var err error
		if lockuptypes.IsSyntheticDenom(gauge.DistributeTo.Denom) {
			gaugeDistributedCoins, err = k.distributeSyntheticInternal(ctx, gauge, filteredLocks, &distrInfo)
		} else {
			gaugeDistributedCoins, err = k.distributeInternal(ctx, gauge, filteredLocks, &distrInfo)
		}
		if err != nil {
			return nil, err
		}
		totalDistributedCoins = totalDistributedCoins.Add(gaugeDistributedCoins...)
	}

	err := k.doDistributionSends(ctx, &distrInfo)
	if err != nil {
		return nil, err
	}
	k.hooks.AfterEpochDistribution(ctx)

	k.checkFinishDistribution(ctx, gauges)
	return totalDistributedCoins, nil
}

// Checks if all non perpetual gauges provided have completed their required distributions.
// If complete, move the gauge from an active to a finished status.
func (k Keeper) checkFinishDistribution(ctx sdk.Context, gauges []types.Gauge) {
	for _, gauge := range gauges {
		// filled epoch is increased in this step and we compare with +1
		// TODO: Wat? we increment filled epochs earlier, this looks wrong and like
		// were not paying out the last epoch of rewards...
		// TODO: @dev explain the above comment further so we can either fix or create issue
		if !gauge.IsPerpetual && gauge.NumEpochsPaidOver <= gauge.FilledEpochs+1 {
			if err := k.moveActiveGaugeToFinishedGauge(ctx, gauge); err != nil {
				panic(err)
			}
		}
	}
}

// Returns sum of coins yet to be distributed for all of the module.
func (k Keeper) GetModuleToDistributeCoins(ctx sdk.Context) sdk.Coins {
	activeGaugesDistr := k.getToDistributeCoinsFromIterator(ctx, k.ActiveGaugesIterator(ctx))
	upcomingGaugesDistr := k.getToDistributeCoinsFromIterator(ctx, k.UpcomingGaugesIteratorAfterTime(ctx, ctx.BlockTime()))
	return activeGaugesDistr.Add(upcomingGaugesDistr...)
}

// GetModuleDistributedCoins returns sum of coins that have been distributed so far for all of the module.
func (k Keeper) GetModuleDistributedCoins(ctx sdk.Context) sdk.Coins {
	activeGaugesDistr := k.getDistributedCoinsFromIterator(ctx, k.ActiveGaugesIterator(ctx))
	finishedGaugesDistr := k.getDistributedCoinsFromIterator(ctx, k.FinishedGaugesIterator(ctx))
	return activeGaugesDistr.Add(finishedGaugesDistr...)
}
