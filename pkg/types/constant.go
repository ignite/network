package types

const (
	// MaxMetadataLength is the max length for metadata attached to chain and campaign
	MaxMetadataLength = 100

	// DefaultUnbondingPeriod is the default unbonding time in seconds
	// 1814400 represents 21 days
	DefaultUnbondingPeriod = 1814400

	// MinimalUnbondingPeriod is the minimal unbonding time that can be set for a chain
	// it must greater than 1 because trusting period for the IBC client is unbonding period - 1
	// and trusting period can't be 0
	MinimalUnbondingPeriod = 2

	// DefaultRevisionHeight is the revision height used by default for creating the monitoring IBC client
	DefaultRevisionHeight = 1
)