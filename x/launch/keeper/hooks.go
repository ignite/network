package keeper

import (
	"context"

	"github.com/ignite/network/x/launch/types"
)

// Implements LaunchHooks interface
var _ types.LaunchHooks = Keeper{}

// RequestCreated calls associated hook if registered
func (k Keeper) RequestCreated(
	ctx context.Context,
	creator string,
	launchID,
	requestID uint64,
	content types.RequestContent,
) error {
	if k.hooks == nil {
		return nil
	}
	return k.hooks.RequestCreated(
		ctx,
		creator,
		launchID,
		requestID,
		content,
	)
}
