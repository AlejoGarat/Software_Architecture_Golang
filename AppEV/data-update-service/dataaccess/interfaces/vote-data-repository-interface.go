package interfaces

import "data-update-service/models/read"

type VoteDataRepository interface {
	UpdateSchedule(read.Vote) error
	UpdateCircuitAgeVotes(read.Vote, read.Voter) error
	UpdateCircuitGenderVotes(read.Vote, read.Voter) error
	UpdateDepartmentData(read.Vote, read.Voter) error
	VoterHasVoted(string, string) bool
	GetVoterByDocument(string, string) (read.Voter, error)
	GetVotes(string) ([]read.VoteGet, error)
	GetCircuits(string) ([]read.Circuit, error)
	GetCandidates(string) ([]read.Candidate, error)
	UpdateTotalVotes(string) error
}
