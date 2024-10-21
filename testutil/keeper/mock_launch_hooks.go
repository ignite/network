package keeper

import (
	launchtypes "github.com/ignite/network/x/launch/types"
)

//go:generate mockery --name LaunchHooks --filename mock_launch_hooks.go --case underscore --output ./mocks
type LaunchHooks interface {
	launchtypes.LaunchHooks
}
