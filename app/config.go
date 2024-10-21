package app

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	networktypes "github.com/ignite/network/pkg/types"
)

func init() {
	// Set prefixes
	accountPubKeyPrefix := networktypes.AccountAddressPrefix + "pub"
	validatorAddressPrefix := networktypes.AccountAddressPrefix + "valoper"
	validatorPubKeyPrefix := networktypes.AccountAddressPrefix + "valoperpub"
	consNodeAddressPrefix := networktypes.AccountAddressPrefix + "valcons"
	consNodePubKeyPrefix := networktypes.AccountAddressPrefix + "valconspub"

	// Set and seal config
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(networktypes.AccountAddressPrefix, accountPubKeyPrefix)
	config.SetBech32PrefixForValidator(validatorAddressPrefix, validatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(consNodeAddressPrefix, consNodePubKeyPrefix)
	config.Seal()
}
