package cli

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	neturl "net/url"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/ignite/network/x/launch/types"
)

const (
	flagMetadataFile      = "metadata-file"
	flagGenesisURL        = "genesis-url"
	flagGenesisConfigFile = "genesis-config"
	flagProjectID         = "project-id"
	flagAccountBalance    = "account-balance"
	flagMetadata          = "metadata"
)

func CmdCreateChain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-chain [genesis-chain-id] [source-url] [source-hash]",
		Short: "Create a new chain for launch",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			argProjectID, err := cmd.Flags().GetInt64(flagProjectID)
			if err != nil {
				return err
			}

			hasProject := false
			projectID := uint64(0)
			if argProjectID >= 0 {
				hasProject = true
				projectID = uint64(argProjectID)
			}

			initialGenesis := types.NewDefaultInitialGenesis()

			// parse genesis url for initialGenesis
			genesisURL, err := cmd.Flags().GetString(flagGenesisURL)
			if err != nil {
				return err
			}
			var genesisHash string
			if genesisURL != "" {
				genesisHash, err = getHashFromURL(cmd.Context(), genesisURL)
				if err != nil {
					return err
				}
				initialGenesis = types.NewGenesisURL(genesisURL, genesisHash)
			}

			// parse genesis config file for initialGenesis
			genesisConfigFile, err := cmd.Flags().GetString(flagGenesisConfigFile)
			if err != nil {
				return err
			}
			if genesisConfigFile != "" {
				initialGenesis = types.NewGenesisConfig(genesisConfigFile)
			}

			// ensure genesisURL and config not being used simultaneously
			if genesisURL != "" && genesisConfigFile != "" {
				return errors.New("cannot use genesisURL and genesis config file")
			}

			var metadataBytes []byte
			metadataFile, _ := cmd.Flags().GetString(flagMetadataFile)
			if metadataFile != "" {
				metadataBytes, err = os.ReadFile(metadataFile)
				if err != nil {
					return err
				}
			} else {
				metadata, _ := cmd.Flags().GetString(flagMetadata)
				metadataBytes = []byte(metadata)
			}

			balanceCoins := sdk.NewCoins()
			balance, err := cmd.Flags().GetString(flagAccountBalance)
			if err != nil {
				return err
			}
			if balance != "" {
				// parse coins argument
				balanceCoins, err = sdk.ParseCoinsNormalized(balance)
				if err != nil {
					return err
				}
			}

			msg := types.NewMsgCreateChain(
				clientCtx.GetFromAddress().String(),
				args[0],
				args[1],
				args[2],
				initialGenesis,
				hasProject,
				projectID,
				balanceCoins,
				metadataBytes,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagGenesisURL, "", "URL for a custom genesis")
	cmd.Flags().String(flagMetadataFile, "", "Set metadata for chain from file content")
	cmd.Flags().String(flagGenesisConfigFile, "", "config file for a custom genesis")
	cmd.Flags().Int64(flagProjectID, -1, "The project id")
	cmd.Flags().String(flagMetadata, "", "Set metadata field for the chain")
	cmd.Flags().String(flagAccountBalance, "", "Set the chain account coin balance")

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// getHashFromURL fetches content from url and returns the hash based on the genesis hash method
func getHashFromURL(ctx context.Context, url string) (string, error) {
	// check url port
	parsedURL, err := neturl.Parse(url)
	if err != nil {
		return "", err
	}
	_, port, err := net.SplitHostPort(parsedURL.Host)
	if err != nil {
		return "", err
	}
	if port != "443" {
		return "", errors.New("port must be 443")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("genesis url fetch error %s", res.Status)
	}
	initialGenesis, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return types.GenesisURLHash(string(initialGenesis)), nil
}
