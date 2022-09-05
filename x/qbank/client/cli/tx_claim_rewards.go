package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/quasarlabs/quasarnode/x/qbank/types"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdClaimRewards() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-rewards [vault-id]",
		Short: "Broadcast message claimRewards",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argVaultID := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgClaimRewards(
				clientCtx.GetFromAddress().String(),
				argVaultID,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
