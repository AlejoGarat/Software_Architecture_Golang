package usecases

import (
	idataaccess "consultant-service/dataaccess/interfaces"
	"consultant-service/helpers"
	"consultant-service/models/read"
	"consultant-service/models/write"
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
		log = write.LoggingModel{Type: "Error", Operation: "Get Department Vote Coverage", Actor: "Consultant", Description: err.Error()}
		departmentUseCase.helpers.LogHelper.SendLog(log)
	}

	return departmentCoverage, err
}
