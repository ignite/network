package constructor

import (
	"cosmossdk.io/core/comet"
	abci "github.com/cometbft/cometbft/abci/types"
)

type Validator struct {
	abci.Validator
}

var _ comet.Validator = (*Validator)(nil)

func (v Validator) Address() []byte {
	return v.Validator.Address
}

func (v Validator) Power() int64 {
	return v.Validator.Power
}

type CommitInfo struct {
	abci.CommitInfo
}

var _ comet.CommitInfo = (*CommitInfo)(nil)

func (e CommitInfo) Round() int32 {
	return e.CommitInfo.Round
}

func (e CommitInfo) Votes() comet.VoteInfos {
	return VoteInfos{e.CommitInfo.Votes}
}

type VoteInfo struct {
	abci.VoteInfo
}

var _ comet.VoteInfo = (*VoteInfo)(nil)

func (v VoteInfo) Validator() comet.Validator {
	return Validator{v.VoteInfo.Validator}
}

func (v VoteInfo) GetBlockIDFlag() comet.BlockIDFlag {
	return comet.BlockIDFlag(v.VoteInfo.BlockIdFlag)
}

type VoteInfos struct {
	votes []abci.VoteInfo
}

var _ comet.VoteInfos = (*VoteInfos)(nil)

func (e VoteInfos) Len() int {
	return len(e.votes)
}

func (e VoteInfos) Get(i int) comet.VoteInfo {
	return VoteInfo{e.votes[i]}
}
