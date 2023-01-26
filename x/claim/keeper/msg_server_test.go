package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/ArkeoNetwork/arkdrop/testutil/keeper"
	"github.com/ArkeoNetwork/arkdrop/x/claim/keeper"
	"github.com/ArkeoNetwork/arkdrop/x/claim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.ClaimKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
