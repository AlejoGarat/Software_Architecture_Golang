package interfaces

import "consultant-service/models/read"

type CircuitUseCase interface {
	GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error)
}
