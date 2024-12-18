package types

import (
	"errors"

	"github.com/ignite/network/pkg/chainid"
)

// Validate checks the chain has valid data
func (m Chain) Validate() error {
	if _, _, err := chainid.ParseGenesisChainID(m.GenesisChainId); err != nil {
		return err
	}

	// A chain that is a mainnet is always associated to a project
	if m.IsMainnet && !m.HasProject {
		return errors.New("chain is a mainnet but not associated to a project")
	}

	// Coins must be valid
	if !m.AccountBalance.IsValid() {
		return errors.New("default account balance sdk.Coins is not valid")
	}

	return nil
}
