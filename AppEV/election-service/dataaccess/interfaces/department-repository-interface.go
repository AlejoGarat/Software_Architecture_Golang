package interfaces

import (
	"election-service/models/read"
)

type DepartmentRepository interface {
	GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error)
}
