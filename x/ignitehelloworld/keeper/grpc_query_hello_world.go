package keeper

import (
	"context"

	"ignite-hello-world/x/ignitehelloworld/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) HelloWorld(goCtx context.Context, req *types.QueryHelloWorldRequest) (*types.QueryHelloWorldResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Process the query
	_ = ctx

	// To respond with a message, use: ignite-hello-worldd q ignitehelloworld hello-world
	return &types.QueryHelloWorldResponse{Message: "Hello, World!!!"}, nil
}
