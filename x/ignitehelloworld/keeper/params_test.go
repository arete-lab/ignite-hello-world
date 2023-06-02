package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "ignite-hello-world/testutil/keeper"
	"ignite-hello-world/x/ignitehelloworld/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.IgnitehelloworldKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
