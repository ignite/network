package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// IsRegistrationEnabled returns true if the current block time is within the allowed registration period
func (k Keeper) IsRegistrationEnabled(ctx context.Context, auctionStartTime time.Time) (bool, error) {
	blockTime := sdk.UnwrapSDKContext(ctx).BlockTime()
	if !blockTime.Before(auctionStartTime) {
		return false, nil
	}

	params, err := k.Params.Get(ctx)
	if err != nil {
		return false, err
	}

	registrationPeriod := params.RegistrationPeriod
	if auctionStartTime.Unix() < int64(registrationPeriod.Seconds()) {
		// subtraction would result in negative value, clamp the result to ~0
		// by making registrationPeriod ~= auctionStartTime
		registrationPeriod = time.Duration(auctionStartTime.Unix()) * time.Second
	}
	// as commented in `Time.Sub()`: To compute t-d for a duration d, use t.Add(-d).
	registrationStart := auctionStartTime.Add(-registrationPeriod)
	return blockTime.After(registrationStart), nil
}
