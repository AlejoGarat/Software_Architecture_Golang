package usecases

import (
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/helpers"
	"election-service/models/read"
	"election-service/models/write"
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
		log = write.LoggingModel{Type: "Error", Operation: "Get Department Vote Coverage", Actor: "Electoral Authority", Description: err.Error()}
		departmentUseCase.helpers.LogHelper.SendLog(log)
	}

	return departmentCoverage, err
}
