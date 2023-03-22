package interfaces

import (
	c "election-service/config"
	"election-service/models/read"
	"election-service/models/write"
)

type ElectionMemoryRepository interface {
	AddElection(write.Election) error
	AddVotersPerDepartment(read.VotersPerDepartment) error
	AddCandidates(candidates []string, electionId string) error
	AddPoliticalParties(politicalParties []string, electionId string) error
	AddDepartments(departments []string, electionId string) error
	GetStartElectionFilters() (c.Configurations, error)
	GetCloseElectionFilters() (c.Configurations, error)
	GetElectionResult(string) (read.ElectionResult, error)
}
