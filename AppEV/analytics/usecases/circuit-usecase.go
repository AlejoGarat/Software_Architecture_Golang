package usecases

import (
	idataaccess "analytics/dataaccess/interfaces"
	"analytics/helpers"
	"analytics/models/read"
	"analytics/models/write"
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
		log = write.LoggingModel{Type: "Error", Operation: "Get Circuit Vote Coverage", Actor: "Consulting Agent", Description: err.Error()}
		circuitUseCase.helpers.LogHelper.SendLog(log)
	}

	return circuitVoteCoverage, err
}
