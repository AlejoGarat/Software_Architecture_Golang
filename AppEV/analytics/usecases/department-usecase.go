package usecases

import (
	idataaccess "analytics/dataaccess/interfaces"
	"analytics/helpers"
	"analytics/models/read"
	"analytics/models/write"
)

type DepartmentUseCase struct {
	departmentRepository idataaccess.DepartmentRepository
	helpers              helpers.Helpers
}

func NewDepartmentUseCase(departmentRepository idataaccess.DepartmentRepository, helpers helpers.Helpers) *DepartmentUseCase {
	return &DepartmentUseCase{departmentRepository: departmentRepository, helpers: helpers}
}

func (departmentUseCase *DepartmentUseCase) GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error) {
	var log write.LoggingModel
	departmentCoverage, err := departmentUseCase.departmentRepository.GetVoteCoveragePerDepartment(electionId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Department Vote Coverage", Actor: "Consulting Agent", Description: err.Error()}
		departmentUseCase.helpers.LogHelper.SendLog(log)
	}

	return departmentCoverage, err
}
