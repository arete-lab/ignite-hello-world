package keeper

import (
	"context"
	"fmt"
	"ignite-hello-world/x/ignitehelloworld/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SayHello(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.QuerySayHelloResponse{Name: fmt.Sprintf("hello %s", req.Name)}, nil
}
