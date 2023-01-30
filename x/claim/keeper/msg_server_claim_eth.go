package keeper

import (
	"context"

	"github.com/ArkeoNetwork/arkdrop/x/claim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ClaimEth(goCtx context.Context, msg *types.MsgClaimEth) (*types.MsgClaimEthResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	// 1. get eth claim
	ethClaim, err := k.GetClaimRecord(ctx, msg.EthAddress, types.ETHEREUM)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get claim record for %s", msg.EthAddress)
	}

	if ethClaim.InitialClaimableAmount.IsZero() {
		return nil, errors.Wrapf(err, "no claimable amount for %s", msg.EthAddress)
	}

	// 2. check if already claimed
	if ethClaim.ActionCompleted[0] {
		return nil, errors.Wrapf(err, "already claimed for %s", msg.EthAddress)
	}
	// 3. validate signature (TODO: implement)

	// set eth claim to completed
	ethClaim.ActionCompleted[types.ForeignChainActionClaim] = true
	err = k.SetClaimRecord(ctx, ethClaim)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set claim record for %s", msg.EthAddress)
	}

	// create new arkeo claim
	arkeoClaim := types.ClaimRecord{
		Address:                msg.Creator,
		Chain:                  types.ARKEO,
		InitialClaimableAmount: ethClaim.InitialClaimableAmount,
		ActionCompleted:        []bool{false, false},
	}
	err = k.SetClaimRecord(ctx, arkeoClaim)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to set claim record for %s", msg.Creator)
	}

	return &types.MsgClaimEthResponse{}, nil
}
