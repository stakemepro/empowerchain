package cli

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/EmpowerPlastic/empowerchain/x/certificates"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        certificates.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", certificates.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(MsgCreateCertificateCmd())
	return cmd
}

func MsgCreateCertificateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-certificate [map of certificate attributes]",
		Short: "create a new certificate",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			currentDir, err := filepath.Abs("./")
			if err == nil {
				fmt.Println("Error: ", err)
			}
			testData := currentDir + args[0]
			fileContent, err := os.ReadFile(testData)
			if err != nil {
				return err
			}

			var data certificates.MsgCreateCertificate
			err = json.Unmarshal(fileContent, &data)
			if err != nil {
				return err
			}
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &data)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
