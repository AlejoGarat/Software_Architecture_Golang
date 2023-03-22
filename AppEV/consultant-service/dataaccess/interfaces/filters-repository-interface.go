package interfaces

import "consultant-service/models/read"

type FilterRepository interface {
	ModifyElectionBeginningFilters(filters read.ElectionBeginningFilters) error
	ModifyElectionEndFilters(filters read.ElectionEndFilters) error
	ModifyVoteIssuanceFilters(filters read.VoteIssuanceFilters) error
}
