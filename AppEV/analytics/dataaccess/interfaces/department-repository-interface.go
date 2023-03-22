package interfaces

import (
	"analytics/models/read"
)

type DepartmentRepository interface {
	GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error)
}
