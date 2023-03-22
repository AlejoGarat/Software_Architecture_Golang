package interfaces

import (
	"votation-service/models/read"
	"votation-service/models/write"
)

type VoteRepository interface {
	AddVote(write.Vote) error
	GetVoterById(string, string) (write.Voter, error)
	GetCandidateById(string, string) (write.Candidate, error)
	GetElectionById(string) (read.Election, error)
	ExistsVoteOfVoter(string, string) bool
	SetConstancyData(write.ConstancyDBData, string) error
	GetConstancyData(read.ConstancyRequest) (write.Constancy, error)
}
