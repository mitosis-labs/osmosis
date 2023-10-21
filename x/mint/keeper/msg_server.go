package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v20/x/mint/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) Mint(goCtx context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// newCoins := sdk.NewCoins(sdk.NewCoin(msg.BaseDenom, sdk.NewIntFromUint64(msg.Amount)))

	_, err := server.Keeper.MintAndTransferCoins(ctx, msg.Account, msg.Amount, msg.BaseDenom, msg.ChannelId, msg.IsIbcDenom)
	if err != nil {
		return nil, err
	}

	return &types.MsgMintResponse{}, nil
}
