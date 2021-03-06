package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
	"github.com/tendermint/go-amino"
	datacmd "gitlab.com/regen-network/regen-ledger/x/data/client/cli"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	dataQueryCmd := &cobra.Command{
		Use:   "data",
		Short: "Querying commands for the data module",
	}

	dataQueryCmd.AddCommand(client.GetCommands(
		datacmd.GetCmdGetData(mc.storeKey, mc.cdc),
	)...)

	dataQueryCmd.AddCommand(client.GetCommands(
		datacmd.GetCmdGetDataBlockHeight(mc.storeKey, mc.cdc),
	)...)

	return dataQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	dataTxCmd := &cobra.Command{
		Use:   "data",
		Short: "Data transactions subcommands",
	}

	dataTxCmd.AddCommand(client.PostCommands(
		datacmd.GetCmdStoreData(mc.cdc),
	)...)

	return dataTxCmd
}
