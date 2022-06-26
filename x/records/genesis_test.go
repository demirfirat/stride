package records_test

import (
	"testing"

	keepertest "github.com/Stride-Labs/stride/testutil/keeper"
	"github.com/Stride-Labs/stride/testutil/nullify"
	"github.com/Stride-Labs/stride/x/records"
	"github.com/Stride-Labs/stride/x/records/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
		DepositRecordList: []types.DepositRecord{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		DepositRecordCount: 2,
	}
	k, ctx := keepertest.RecordsKeeper(t)
	records.InitGenesis(ctx, *k, genesisState)
	got := records.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.DepositRecordList, got.DepositRecordList)
	require.Equal(t, genesisState.DepositRecordCount, got.DepositRecordCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
