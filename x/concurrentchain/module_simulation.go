package concurrentchain

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"github.com/totegamma/concurrent-chain/testutil/sample"
	concurrentchainsimulation "github.com/totegamma/concurrent-chain/x/concurrentchain/simulation"
	"github.com/totegamma/concurrent-chain/x/concurrentchain/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = concurrentchainsimulation.FindAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
	_ = rand.Rand{}
)

const (
	opWeightMsgCreateAddressbook = "op_weight_msg_addressbook"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateAddressbook int = 100

	opWeightMsgUpdateAddressbook = "op_weight_msg_addressbook"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUpdateAddressbook int = 100

	opWeightMsgDeleteAddressbook = "op_weight_msg_addressbook"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDeleteAddressbook int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module.
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	concurrentchainGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		AddressbookList: []types.Addressbook{
			{
				Creator: sample.AccAddress(),
				Index:   "0",
			},
			{
				Creator: sample.AccAddress(),
				Index:   "1",
			},
		},
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&concurrentchainGenesis)
}

// RegisterStoreDecoder registers a decoder.
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// ProposalContents doesn't return any content functions for governance proposals.
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgCreateAddressbook int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateAddressbook, &weightMsgCreateAddressbook, nil,
		func(_ *rand.Rand) {
			weightMsgCreateAddressbook = defaultWeightMsgCreateAddressbook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateAddressbook,
		concurrentchainsimulation.SimulateMsgCreateAddressbook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUpdateAddressbook int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUpdateAddressbook, &weightMsgUpdateAddressbook, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateAddressbook = defaultWeightMsgUpdateAddressbook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUpdateAddressbook,
		concurrentchainsimulation.SimulateMsgUpdateAddressbook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDeleteAddressbook int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDeleteAddressbook, &weightMsgDeleteAddressbook, nil,
		func(_ *rand.Rand) {
			weightMsgDeleteAddressbook = defaultWeightMsgDeleteAddressbook
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDeleteAddressbook,
		concurrentchainsimulation.SimulateMsgDeleteAddressbook(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}

// ProposalMsgs returns msgs used for governance proposals for simulations.
func (am AppModule) ProposalMsgs(simState module.SimulationState) []simtypes.WeightedProposalMsg {
	return []simtypes.WeightedProposalMsg{
		simulation.NewWeightedProposalMsg(
			opWeightMsgCreateAddressbook,
			defaultWeightMsgCreateAddressbook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				concurrentchainsimulation.SimulateMsgCreateAddressbook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgUpdateAddressbook,
			defaultWeightMsgUpdateAddressbook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				concurrentchainsimulation.SimulateMsgUpdateAddressbook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		simulation.NewWeightedProposalMsg(
			opWeightMsgDeleteAddressbook,
			defaultWeightMsgDeleteAddressbook,
			func(r *rand.Rand, ctx sdk.Context, accs []simtypes.Account) sdk.Msg {
				concurrentchainsimulation.SimulateMsgDeleteAddressbook(am.accountKeeper, am.bankKeeper, am.keeper)
				return nil
			},
		),
		// this line is used by starport scaffolding # simapp/module/OpMsg
	}
}
