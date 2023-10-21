package tx

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/osmosis-labs/osmosis/v20/x/mint/types"
)

var _ = strconv.Itoa(0)

func CmdMint() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint <account> <amount> <denom> <channel-id> <is-ibc-denom>",
		Short: "Send mint",
		Long: `Send mint.{{.ExampleHeader}}
{{.CommandPrefix}} mint mito11vmx8jtggpd9u7qr0t8vxclycz85u925sazglr7 umito 1000`,
		Args: cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			account := args[0]

			amount, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			baseDenom := args[2]

			channelId := args[3]

			isIbcDenom, err := strconv.ParseBool(args[4])
			if err != nil {
				return err
			}

			msg := types.NewMsgMint(
				clientCtx.GetFromAddress().String(),
				account,
				amount,
				baseDenom,
				channelId,
				isIbcDenom,
			)
			if err = msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "send query all balances")

	return cmd
}
