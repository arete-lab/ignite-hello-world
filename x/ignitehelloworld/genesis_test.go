package ignitehelloworld_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "ignite-hello-world/testutil/keeper"
	"ignite-hello-world/testutil/nullify"
	"ignite-hello-world/x/ignitehelloworld"
	"ignite-hello-world/x/ignitehelloworld/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.IgnitehelloworldKeeper(t)
	ignitehelloworld.InitGenesis(ctx, *k, genesisState)
	got := ignitehelloworld.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
