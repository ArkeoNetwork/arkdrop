package keeper_test

import (
	"crypto/ecdsa"
	"strings"
	"testing"

	"github.com/ArkeoNetwork/arkdrop/x/claim/keeper"
	"github.com/ArkeoNetwork/arkdrop/x/claim/types"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
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

func TestIsValidClaimSignature(t *testing.T) {
	// generate a random eth address
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	addressEth := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	addrArkeo := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()

	message, err := keeper.GenerateClaimTypedDataBytes(addressEth, addrArkeo, "5000")
	require.NoError(t, err)
	hash := crypto.Keccak256(message)
	signature, err := crypto.Sign(hash, privateKey)
	require.NoError(t, err)

	sigString := hexutil.Encode(signature)

	// check if signature is valid
	valid, err := keeper.IsValidClaimSignature(strings.ToLower(addressEth), addrArkeo, "5000", sigString)
	require.NoError(t, err)
	require.True(t, valid)

	// if we modify the message, signature should be invalid
	_, err = keeper.IsValidClaimSignature(addressEth, addrArkeo, "5001", sigString)
	require.Error(t, err)

	// if we modify the arkeo address, signature should be invalid
	addrArkeo2 := sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
	_, err = keeper.IsValidClaimSignature(addressEth, addrArkeo2, "5000", sigString)
	require.Error(t, err)

	// if we modify the eth address, signature should be invalid
	_, err = keeper.IsValidClaimSignature("0xbd3afb0bb76683ecb4225f9dbc91f998713c3b01", addrArkeo, "5000", sigString)
	require.Error(t, err)
}
