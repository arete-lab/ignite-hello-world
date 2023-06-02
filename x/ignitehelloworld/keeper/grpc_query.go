package keeper

import (
	"ignite-hello-world/x/ignitehelloworld/types"
)

var _ types.QueryServer = Keeper{}
