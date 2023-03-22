package interfaces

import (
	"consultant-service/models/read"
)

type CircuitRepository interface {
	GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error)
}
