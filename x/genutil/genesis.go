package genutil

//nolint:goimports
import (
	"encoding/json"
	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/auth"
	abci "github.com/tendermint/tendermint/abci/types"

	"git.dsr-corporation.com/zb-ledger/zb-ledger/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis - initialize accounts and deliver genesis transactions.
func InitGenesis(ctx sdk.Context, cdc *codec.Codec, authKeeper types.AuthKeeper, validatorKeeper types.ValidatorKeeper,
	deliverTx deliverTxfn, genesisState GenesisState) []abci.ValidatorUpdate {
	// load the accounts
	for _, acc := range genesisState.Accounts {
		acc = authKeeper.NewAccountWithNumber(ctx, acc) // set account number
		authKeeper.SetAccount(ctx, acc)
	}

	// deliver validator transactions
	var validators []abci.ValidatorUpdate
	if len(genesisState.GenTxs) > 0 {
		validators = DeliverGenTxs(ctx, cdc, genesisState.GenTxs, validatorKeeper, deliverTx)
	}

	return validators
}

func IterateGenesisAccounts(cdc *codec.Codec, appGenesis map[string]json.RawMessage,
	iterateFn func(auth.Account) (stop bool)) {
	genesisState := types.GetGenesisStateFromAppState(cdc, appGenesis)
	for _, acc := range genesisState.Accounts {
		if iterateFn(acc) {
			break
		}
	}
}
