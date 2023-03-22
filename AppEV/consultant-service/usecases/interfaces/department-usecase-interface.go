package interfaces

import "consultant-service/models/read"

type DepartmentUseCase interface {
	GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error)
}
