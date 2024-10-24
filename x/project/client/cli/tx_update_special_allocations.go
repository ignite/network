package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/ignite/network/x/project/types"
)

func CmdUpdateSpecialAllocations() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-special-allocations [project-id] [genesis-distribution] [claimable-airdrop]",
		Short: "Update special allocations for the project",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argProjectID, err := cast.ToUint64E(args[0])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			genesisDistribution, err := types.NewShares(args[1])
			if err != nil {
				return err
			}

			claimableAirdrop, err := types.NewShares(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateSpecialAllocations(
				clientCtx.GetFromAddress().String(),
				argProjectID,
				types.NewSpecialAllocations(genesisDistribution, claimableAirdrop),
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
