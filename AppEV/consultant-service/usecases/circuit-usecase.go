package usecases

import (
	idataaccess "consultant-service/dataaccess/interfaces"
	"consultant-service/helpers"
	"consultant-service/models/read"
	"consultant-service/models/write"
)

type CircuitUseCase struct {
	circuitRepository idataaccess.CircuitRepository
	helpers           helpers.Helpers
}

func NewCircuitUseCase(circuitRepository idataaccess.CircuitRepository, helpers helpers.Helpers) *CircuitUseCase {
	return &CircuitUseCase{circuitRepository: circuitRepository, helpers: helpers}
}

func (circuitUseCase *CircuitUseCase) GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error) {
	var log write.LoggingModel
	circuitVoteCoverage, err := circuitUseCase.circuitRepository.GetVoteCoveragePerCircuit(electionId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Circuit Vote Coverage", Actor: "Consultant", Description: err.Error()}
		circuitUseCase.helpers.LogHelper.SendLog(log)
	}

	return circuitVoteCoverage, err
}
