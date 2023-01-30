package keeper_test

import (
	"testing"

	"github.com/ArkeoNetwork/arkdrop/x/claim/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestClaimEth(t *testing.T) {
	msgServer, keeper, ctx := setupMsgServer(t)
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// create valid eth claimrecords
	addrEth := "0xDAFEA492D9c6733ae3d56b7Ed1ADB60692c98Bc5" // random invalid eth address
	addrArkeo := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()

	claimRecord := types.ClaimRecord{
		Chain:                  types.ETHEREUM,
		Address:                addrEth,
		InitialClaimableAmount: sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 100)),
		ActionCompleted:        []bool{false, false},
	}
	err := keeper.SetClaimRecord(sdkCtx, claimRecord)
	require.NoError(t, err)

	claimMessage := types.MsgClaimEth{
		Creator:    addrArkeo,
		EthAddress: addrEth,
		Signature:  "",
	}

	_, err = msgServer.ClaimEth(ctx, &claimMessage)
	require.NoError(t, err)

	// check if claimrecord is updated
	claimRecord, err = keeper.GetClaimRecord(sdkCtx, addrEth, types.ETHEREUM)
	require.NoError(t, err)
	require.True(t, claimRecord.ActionCompleted[types.ForeignChainActionClaim])

	// confirm we have a claimrecord for arkeo
	claimRecord, err = keeper.GetClaimRecord(sdkCtx, addrArkeo, types.ARKEO)
	require.NoError(t, err)
	require.Equal(t, claimRecord.Address, addrArkeo)
	require.Equal(t, claimRecord.Chain, types.ARKEO)
	require.Equal(t, claimRecord.InitialClaimableAmount, sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 100)))
	require.False(t, claimRecord.ActionCompleted[types.ActionVote])
	require.False(t, claimRecord.ActionCompleted[types.ActionDelegateStake])
}
