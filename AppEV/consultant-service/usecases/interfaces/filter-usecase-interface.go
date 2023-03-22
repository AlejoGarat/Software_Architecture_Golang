package interfaces

import (
	"consultant-service/models/read"
)

type FilterUseCase interface {
	ModifyElectionBeginningFilters(read.ElectionBeginningFilters) error
	ModifyElectionEndFilters(read.ElectionEndFilters) error
	ModifyVoteIssuanceFilters(read.VoteIssuanceFilters) error
}
