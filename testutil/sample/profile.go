package sample

import (
	"math/rand"

	profile "github.com/ignite/network/x/profile/types"
)

// MsgCreateCoordinator returns a sample MsgCreateCoordinator
func MsgCreateCoordinator(coordAddress string) profile.MsgCreateCoordinator {
	return *profile.NewMsgCreateCoordinator(
		coordAddress,
		coordAddress,
		"https://cosmos.network/"+coordAddress,
		coordAddress+" details",
	)
}

// MsgDisableCoordinator returns a sample MsgDisableCoordinator
func MsgDisableCoordinator(coordAddress string) profile.MsgDisableCoordinator {
	return *profile.NewMsgDisableCoordinator(
		coordAddress,
	)
}

// MsgUpdateCoordinatorDescription returns a sample MsgUpdateCoordinatorDescription
func MsgUpdateCoordinatorDescription(coordAddress string) profile.MsgUpdateCoordinatorDescription {
	return *profile.NewMsgUpdateCoordinatorDescription(
		coordAddress,
		coordAddress+" update identity",
		"https://cosmos.network/update/"+coordAddress,
		coordAddress+" update details",
	)
}

// ValidatorDescription returns a sample ValidatorDescription
func ValidatorDescription(desc string) profile.ValidatorDescription {
	return profile.ValidatorDescription{
		Identity:        desc,
		Moniker:         "moniker " + desc,
		Website:         "https://cosmos.network/" + desc,
		SecurityContact: "foo",
		Details:         desc + " details",
	}
}

// Coordinator returns a sample Coordinator
func Coordinator(r *rand.Rand, address string) profile.Coordinator {
	return profile.Coordinator{
		Address:     address,
		Description: CoordinatorDescription(r),
		Active:      true,
	}
}

// CoordinatorDescription returns a sample CoordinatorDescription
func CoordinatorDescription(r *rand.Rand) profile.CoordinatorDescription {
	return profile.CoordinatorDescription{
		Identity: String(r, 10),
		Website:  String(r, 10),
		Details:  String(r, 30),
	}
}

// ProfileGenesisState returns a sample genesis state for the profile module
func ProfileGenesisState(r *rand.Rand, addresses ...string) profile.GenesisState {
	for len(addresses) < 7 {
		addresses = append(addresses, Address(r))
	}
	operatorAddresses := []string{OperatorAddress(r), OperatorAddress(r)}
	return profile.GenesisState{
		CoordinatorList: []profile.Coordinator{
			{
				CoordinatorId: 0,
				Address:       addresses[0],
				Description:   CoordinatorDescription(r),
				Active:        true,
			},
			{
				CoordinatorId: 1,
				Address:       addresses[1],
				Description:   CoordinatorDescription(r),
				Active:        true,
			},
			{
				CoordinatorId: 2,
				Address:       addresses[2],
				Description:   CoordinatorDescription(r),
				Active:        true,
			},
			{
				CoordinatorId: 3,
				Address:       addresses[3],
				Description:   CoordinatorDescription(r),
				Active:        true,
			},
			{
				CoordinatorId: 4,
				Address:       addresses[4],
				Description:   CoordinatorDescription(r),
				Active:        true,
			},
		},
		CoordinatorByAddress: []profile.CoordinatorByAddress{
			{
				Address:       addresses[0],
				CoordinatorId: 0,
			},
			{
				Address:       addresses[1],
				CoordinatorId: 1,
			},
			{
				Address:       addresses[2],
				CoordinatorId: 2,
			},
			{
				Address:       addresses[3],
				CoordinatorId: 3,
			},
			{
				Address:       addresses[4],
				CoordinatorId: 4,
			},
		},
		CoordinatorCount: 5,
		ValidatorList: []profile.Validator{
			{
				Address:           addresses[5],
				Description:       ValidatorDescription(String(r, 10)),
				OperatorAddresses: []string{operatorAddresses[0]},
			},
			{
				Address:           addresses[6],
				Description:       ValidatorDescription(String(r, 10)),
				OperatorAddresses: []string{operatorAddresses[1]},
			},
		},
		ValidatorsByOperatorAddress: []profile.ValidatorByOperatorAddress{
			{
				OperatorAddress:  operatorAddresses[0],
				ValidatorAddress: addresses[5],
			},
			{
				OperatorAddress:  operatorAddresses[1],
				ValidatorAddress: addresses[6],
			},
		},
	}
}
