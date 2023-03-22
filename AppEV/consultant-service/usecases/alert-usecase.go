package usecases

import (
	idataaccess "consultant-service/dataaccess/interfaces"
	"consultant-service/helpers"
	"consultant-service/models/write"
	"errors"
)

type AlertUseCase struct {
	alertRepository idataaccess.AlertRepository
	helpers         helpers.Helpers
}

func NewAlertUseCase(alertRepository idataaccess.AlertRepository, helpers helpers.Helpers) *AlertUseCase {
	return &AlertUseCase{alertRepository: alertRepository, helpers: helpers}
}

func (alertUseCase *AlertUseCase) ModifyAlertConfiguration(alertConfig write.AlertConfiguration) error {

	if alertConfig.ElectionId == "" {
		return errors.New("Config must have an election id.")
	}

	var log write.LoggingModel
	err := alertUseCase.alertRepository.ModifyAlertConfiguration(alertConfig)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Modify Alert Configuration", Actor: "Consultant", Description: err.Error()}
		alertUseCase.helpers.LogHelper.SendLog(log)
	}

	return err
}

func (alertUseCase *AlertUseCase) GetAlertConfiguration(electionId string) (write.AlertConfiguration, error) {
	var log write.LoggingModel
	alertConfig, err := alertUseCase.alertRepository.GetAlertConfiguration(electionId)

	if err != nil {
		log = write.LoggingModel{Type: "Error", Operation: "Get Alert Configuration", Actor: "Consultant", Description: err.Error()}
		alertUseCase.helpers.LogHelper.SendLog(log)
	}

	return alertConfig, err
}
