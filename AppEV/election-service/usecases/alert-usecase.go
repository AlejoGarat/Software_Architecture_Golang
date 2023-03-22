package usecases

import (
	idataaccess "election-service/dataaccess/interfaces"
	"election-service/helpers"
	"election-service/models/read"
	"election-service/models/write"
)

type AlertUseCase struct {
	alertRepository idataaccess.AlertRepository
	helpers         helpers.Helpers
}

func NewAlertUseCase(alertRepository idataaccess.AlertRepository, helpers helpers.Helpers) *AlertUseCase {
	return &AlertUseCase{alertRepository: alertRepository, helpers: helpers}
}

func (alertUseCase *AlertUseCase) GetAlertConfiguration(electionId string) (read.AlertConfiguration, error) {
	var log write.LoggingModel
	alertConfig, err := alertUseCase.alertRepository.GetAlertConfiguration(electionId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Alert Configuration", Actor: "Electoral Authority", Description: err.Error()}
		alertUseCase.helpers.LogHelper.SendLog(log)
	}

	return alertConfig, err
}
