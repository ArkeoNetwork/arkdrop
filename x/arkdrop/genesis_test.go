package arkdrop_test

import (
	"testing"

	keepertest "github.com/ArkeoNetwork/arkdrop/testutil/keeper"
	"github.com/ArkeoNetwork/arkdrop/testutil/nullify"
	"github.com/ArkeoNetwork/arkdrop/x/arkdrop"
	"github.com/ArkeoNetwork/arkdrop/x/arkdrop/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ArkdropKeeper(t)
	arkdrop.InitGenesis(ctx, *k, genesisState)
	got := arkdrop.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
