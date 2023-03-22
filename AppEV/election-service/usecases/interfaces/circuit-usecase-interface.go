package interfaces

import "election-service/models/read"

type CircuitUseCase interface {
	GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error)
}
