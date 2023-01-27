package keeper_test

import (
	"testing"
	"time"

	testkeeper "github.com/ArkeoNetwork/arkdrop/testutil/keeper"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"

	"github.com/ArkeoNetwork/arkdrop/x/claim/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestGetClaimRecord(t *testing.T) {
	keeper, ctx := testkeeper.ClaimKeeper(t)
	airdropStartTime := time.Now()
	params := types.Params{
		AirdropStartTime:   airdropStartTime,
		DurationUntilDecay: types.DefaultDurationUntilDecay,
		DurationOfDecay:    types.DefaultDurationOfDecay,
		ClaimDenom:         types.DefaultClaimDenom,
	}
	keeper.SetParams(ctx, params)

	addr1 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	addr2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())
	addr3 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address())

	claimRecords := []types.ClaimRecord{
		{
			Address:                addr1.String(),
			InitialClaimableAmount: sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 100)),
			ActionCompleted:        []bool{false, false},
		},
		{
			Address:                addr2.String(),
			InitialClaimableAmount: sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 200)),
			ActionCompleted:        []bool{false, false},
		},
	}
	err := keeper.SetClaimRecords(ctx, claimRecords)
	require.NoError(t, err)

	coins1, err := keeper.GetUserTotalClaimable(ctx, addr1)
	require.NoError(t, err)
	require.Equal(t, coins1, claimRecords[0].InitialClaimableAmount, coins1.String())

	coins2, err := keeper.GetUserTotalClaimable(ctx, addr2)
	require.NoError(t, err)
	require.Equal(t, coins2, claimRecords[1].InitialClaimableAmount)

	coins3, err := keeper.GetUserTotalClaimable(ctx, addr3)
	require.NoError(t, err)
	require.Equal(t, coins3, sdk.Coins{})

	// get rewards amount per action
	coins4, err := keeper.GetClaimableAmountForAction(ctx, addr1, types.ActionVote)
	require.NoError(t, err)
	require.Equal(t, coins4.String(), sdk.NewCoins(sdk.NewInt64Coin(types.DefaultClaimDenom, 50)).String())

	// get completed activities
	claimRecord, err := keeper.GetClaimRecord(ctx, addr1)
	require.NoError(t, err)
	for i := range types.Action_name {
		require.False(t, claimRecord.ActionCompleted[i])
	}
}
