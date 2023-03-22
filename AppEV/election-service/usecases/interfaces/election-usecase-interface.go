package interfaces

import "election-service/models/read"

type ElectionUsecase interface {
	AddElection(election read.ElectionData) error
	GetElectionResult(string) (read.ElectionResult, error)
}
