package interfaces

import (
	"consultant-service/models/read"
)

type DepartmentRepository interface {
	GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error)
}
