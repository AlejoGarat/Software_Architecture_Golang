package interfaces

import (
	"analytics/models/read"
)

type CircuitRepository interface {
	GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error)
}
