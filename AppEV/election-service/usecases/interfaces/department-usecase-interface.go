package interfaces

import "election-service/models/read"

type DepartmentUseCase interface {
	GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error)
}
