package types

import "context"

// Event Hooks
// These can be utilized to communicate between a fundraising keeper and other keepers.
// The other keepers must implement this interface, which then the fundraising keeper can call.

// LaunchHooks event hooks for launch module
type LaunchHooks interface {
	RequestCreated(
		ctx context.Context,
		creator string,
		launchID,
		requestID uint64,
		content RequestContent,
	) error
}

// MultiLaunchHooks combines multiple launch hooks
type MultiLaunchHooks []LaunchHooks

func NewMultiLaunchHooks(hooks ...LaunchHooks) MultiLaunchHooks {
	return hooks
}

func (h MultiLaunchHooks) RequestCreated(
	ctx context.Context,
	creator string,
	launchID,
	requestID uint64,
	content RequestContent,
) error {
	for i := range h {
		if err := h[i].RequestCreated(
			ctx,
			creator,
			launchID,
			requestID,
			content,
		); err != nil {
			return err
		}
	}
	return nil
}
