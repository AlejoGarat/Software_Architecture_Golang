package interfaces

import (
	"election-service/models/read"
	"election-service/models/write"
)

type ElectionRepository interface {
	AddElection(write.Election) error
	AddVoter(read.Voter, string) error
	AddCandiadtes([]write.Candidate, string) error
	AddPoliticalParties([]write.PoliticalParty, string) error
	AddCircuits([]write.Circuit) error
	GetVotes(string) ([]read.Vote, error)
	GetCandidate(string) (read.Candidate, error)
	AddVotersPerAge(map[string]map[int]int, string) error
	AddVotersPerGender(map[string]map[string]int, string) error
	AddVotersPerDepartmentByGenderAndAge(map[string]map[string]int, map[string]map[int]int, string) error
}
