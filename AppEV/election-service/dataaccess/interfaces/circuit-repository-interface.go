package interfaces

import (
	"election-service/models/read"
)

type CircuitRepository interface {
	GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error)
}
