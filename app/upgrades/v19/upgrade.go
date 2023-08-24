package v19

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	gammtypes "github.com/osmosis-labs/osmosis/v17/x/gamm/types"
	epochtypes "github.com/osmosis-labs/osmosis/x/epochs/types"

	"github.com/osmosis-labs/osmosis/v17/app/keepers"
	"github.com/osmosis-labs/osmosis/v17/app/upgrades"
)

// OSMO / DAI CL pool ID
const firstCLPoolId = 1066

func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	bpm upgrades.BaseAppParamManager,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, plan upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		// Run migrations before applying any other state changes.
		// NOTE: DO NOT PUT ANY STATE CHANGES BEFORE RunMigrations().
		migrations, err := mm.RunMigrations(ctx, configurator, fromVM)
		if err != nil {
			return nil, err
		}

		for id := 1; id < firstCLPoolId; id++ {
			resetSumtree(keepers, ctx, uint64(id))
		}

		epochs := keepers.EpochsKeeper.AllEpochInfos(ctx)
		desiredEpochInfo := epochtypes.EpochInfo{}
		for _, epoch := range epochs {
			if epoch.Identifier == "day" {
				epoch.Duration = time.Minute * 15
				epoch.CurrentEpochStartTime = time.Now().Add(-epoch.Duration).Add(time.Minute)
				desiredEpochInfo = epoch
				keepers.EpochsKeeper.DeleteEpochInfo(ctx, epoch.Identifier)
			}
		}
		keepers.EpochsKeeper.SetEpochInfo(ctx, desiredEpochInfo)

		return migrations, nil
	}
}

func resetSumtree(keepers *keepers.AppKeepers, ctx sdk.Context, id uint64) {
	denom := gammtypes.GetPoolShareDenom(id)
	keepers.LockupKeeper.RebuildAccumulationStoreForDenom(ctx, denom)
}