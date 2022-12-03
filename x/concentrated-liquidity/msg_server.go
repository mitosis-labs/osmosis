package concentrated_liquidity

import (
	"context"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"

	clmodel "github.com/osmosis-labs/osmosis/v13/x/concentrated-liquidity/model"
	"github.com/osmosis-labs/osmosis/v13/x/concentrated-liquidity/types"
)

type msgServer struct {
	keeper *Keeper
}

func NewMsgServerImpl(keeper *Keeper) clmodel.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

// TODO: tests, including events
func (server msgServer) CreatePosition(goCtx context.Context, msg *clmodel.MsgCreatePosition) (*clmodel.MsgCreatePositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	actualAmount0, actualAmount1, liquidityCreated, err := server.keeper.createPosition(ctx, msg.PoolId, sender, msg.TokenDesired0.Amount, msg.TokenDesired1.Amount, msg.TokenMinAmount0, msg.TokenMinAmount1, msg.LowerTick, msg.UpperTick)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.TypeEvtWithdrawPosition,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeAmount0, actualAmount0.String()),
			sdk.NewAttribute(types.AttributeAmount1, actualAmount1.String()),
			sdk.NewAttribute(types.AttributeLiquidity, liquidityCreated.String()),
			sdk.NewAttribute(types.AttributeLowerTick, strconv.FormatInt(msg.LowerTick, 10)),
			sdk.NewAttribute(types.AttributeUpperTick, strconv.FormatInt(msg.UpperTick, 10)),
		),
	})

	return &clmodel.MsgCreatePositionResponse{Amount0: actualAmount0, Amount1: actualAmount1}, nil
}

// TODO: tests, including events
func (server msgServer) WithdrawPosition(goCtx context.Context, msg *clmodel.MsgWithdrawPosition) (*clmodel.MsgWithdrawPositionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	amount0, amount1, err := server.keeper.withdrawPosition(ctx, msg.PoolId, sender, msg.LowerTick, msg.UpperTick, msg.LiquidityAmount.ToDec())
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
		),
		sdk.NewEvent(
			types.TypeEvtWithdrawPosition,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.AttributeValueCategory),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Sender),
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.FormatUint(msg.PoolId, 10)),
			sdk.NewAttribute(types.AttributeLiquidity, msg.LiquidityAmount.String()),
			sdk.NewAttribute(types.AttributeAmount0, amount0.String()),
			sdk.NewAttribute(types.AttributeAmount1, amount1.String()),
			sdk.NewAttribute(types.AttributeLowerTick, strconv.FormatInt(msg.LowerTick, 10)),
			sdk.NewAttribute(types.AttributeUpperTick, strconv.FormatInt(msg.UpperTick, 10)),
		),
	})

	return &clmodel.MsgWithdrawPositionResponse{Amount0: amount0, Amount1: amount1}, nil
}
