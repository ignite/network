package types

import (
	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	networktypes "github.com/ignite/network/pkg/types"
)

func NewMsgCreateClient(
	creator string,
	launchID uint64,
	consensusState networktypes.ConsensusState,
	validatorSet networktypes.ValidatorSet,
	unbondingPeriod int64,
	revisionHeight uint64,
) *MsgCreateClient {
	return &MsgCreateClient{
		Creator:         creator,
		LaunchID:        launchID,
		ConsensusState:  consensusState,
		ValidatorSet:    validatorSet,
		UnbondingPeriod: unbondingPeriod,
		RevisionHeight:  revisionHeight,
	}
}

func (msg MsgCreateClient) Type() string {
	return sdk.MsgTypeURL(&MsgCreateClient{})
}

func (msg *MsgCreateClient) ValidateBasic() error {
	// validate consensus state
	tmConsensusState, err := msg.ConsensusState.ToTendermintConsensusState()
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidConsensusState, err.Error())
	}
	if err := tmConsensusState.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(ErrInvalidConsensusState, err.Error())
	}

	// validate validator set
	tmValidatorSet, err := msg.ValidatorSet.ToTendermintValidatorSet()
	if err != nil {
		return sdkerrors.Wrap(ErrInvalidValidatorSet, err.Error())
	}
	if err := tmValidatorSet.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(ErrInvalidValidatorSet, err.Error())
	}

	// check validator set hash matches consensus state
	if !networktypes.CheckValidatorSetHash(tmValidatorSet, tmConsensusState) {
		return sdkerrors.Wrap(ErrInvalidValidatorSetHash, "validator set hash doesn't match the consensus state")
	}

	// unbonding period must greater than 1 because trusting period for the IBC client is unbonding period - 1
	// and trusting period can't be 0
	if msg.UnbondingPeriod < networktypes.MinimalUnbondingPeriod {
		return sdkerrors.Wrap(ErrInvalidUnbondingPeriod, "unbonding period must be greater than 1")
	}

	// check revision height is non-null
	if msg.RevisionHeight == 0 {
		return sdkerrors.Wrap(ErrInvalidRevisionHeight, "revision height must be non-null")
	}

	return nil
}
