package interfaces

import "analytics/models/read"

type DepartmentUseCase interface {
	GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error)
}
