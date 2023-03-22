package interfaces

import "analytics/models/read"

type CircuitUseCase interface {
	GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error)
}
