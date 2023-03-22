package interfaces

import "data-update-service/models/read"

type MemoryRepository interface {
	UpdateVotesPerParty([]read.StringIntParty, string) error
	UpdateVotesPerCandidate([]read.StringIntCandidate, string) error
	UpdateVotesPerDepartment([]read.StringIntDepartment, string) error
	UpdateTotalVotes(string) error
}
