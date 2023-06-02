package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "ignite-hello-world/testutil/keeper"
	"ignite-hello-world/x/ignitehelloworld/keeper"
	"ignite-hello-world/x/ignitehelloworld/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.IgnitehelloworldKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
