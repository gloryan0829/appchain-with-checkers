package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/alice/checkers/testutil/keeper"
	"github.com/alice/checkers/x/checkers"
	"github.com/alice/checkers/x/checkers/keeper"
	"github.com/alice/checkers/x/checkers/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestRejectGameByRedOneMoveRemovedGame(t *testing.T) {
    msgServer, keeper, context := setupMsgServerWithOneGameForRejectGame(t)
    msgServer.PlayMove(context, &types.MsgPlayMove{
        Creator:   bob,
        GameIndex: "1",
        FromX:     1,
        FromY:     2,
        ToX:       2,
        ToY:       3,
    })
    msgServer.RejectGame(context, &types.MsgRejectGame{
        Creator:   carol,
        GameIndex: "1",
    })
    systemInfo, found := keeper.GetSystemInfo(sdk.UnwrapSDKContext(context))
    require.True(t, found)
    require.EqualValues(t, types.SystemInfo{
        NextId: 2,
    }, systemInfo)
    _, found = keeper.GetStoredGame(sdk.UnwrapSDKContext(context), "1")
    require.False(t, found)
}

func setupMsgServerWithOneGameForRejectGame(t testing.TB) (types.MsgServer, keeper.Keeper, context.Context) {
	k, ctx := keepertest.CheckersKeeper(t)
	checkers.InitGenesis(ctx, *k, *types.DefaultGenesis())
	server := keeper.NewMsgServerImpl(*k)
	context := sdk.WrapSDKContext(ctx)
	server.CreateGame(context, &types.MsgCreateGame{
		Creator: alice,
		Black:   bob,
		Red:     carol,
	})
	return server, *k, context
}
