package usecases

import (
	"fmt"
	idataaccess "monitoring-service/dataaccess/interfaces"
	"monitoring-service/rabbit/workers"
	"strconv"
)

type ConfigurationUseCase struct {
	configurationRepository idataaccess.ConfigurationRepository
}

func NewConfigurationUseCase(configurationRepository idataaccess.ConfigurationRepository) *ConfigurationUseCase {
	return &ConfigurationUseCase{
		configurationRepository: configurationRepository,
	}
}

func (configurationUseCase *ConfigurationUseCase) AnalyzeValues(worker workers.Worker) ([]string, error) {
	var messagesToSend []string

	elections, err := configurationUseCase.configurationRepository.GetElections()

	if err != nil {
		return messagesToSend, err
	}

	for _, election := range elections {

		alertConfig, err := configurationUseCase.configurationRepository.GetAlertConfiguration(election.ElectionId)

		if err != nil {
			return messagesToSend, err
		}

		realMaxVoteAmount, err := configurationUseCase.configurationRepository.GetRealMaxVoteAmountValue(election.ElectionId)

		if err != nil {
			return messagesToSend, err
		}

		realMaxConstancyAmount, err := configurationUseCase.configurationRepository.GetRealMaxConstancyAmountValue(election.ElectionId)

		if err != nil {
			fmt.Print(err.Error())
			return messagesToSend, err
		}

		if realMaxVoteAmount > alertConfig.MaxVoteAmount {
			message := "WARNING: Max Vote Amount value has suffered a deviation: " + "EXPECTED: " + strconv.Itoa(alertConfig.MaxVoteAmount) + " ACTUAL: " + strconv.Itoa(realMaxVoteAmount) + "|||"
			messagesToSend = append(messagesToSend, message)
		}

		if realMaxConstancyAmount > alertConfig.MaxConstancyAmount {
			message := "WARNING: Max Constancy Amount value has suffered a deviation: " + "EXPECTED: " + strconv.Itoa(alertConfig.MaxConstancyAmount) + " ACTUAL: " + strconv.Itoa(realMaxConstancyAmount)
			messagesToSend = append(messagesToSend, message)
		}
	}

	return messagesToSend, nil
}

/*func (configurationUseCase *ConfigurationUseCase) GetAlertConfiguration() (models.ConfigurationModel, error) {
	return configurationUseCase.configurationRepository.GetAlertConfiguration()
}

func (configurationUseCase *ConfigurationUseCase) GetRealMaxVoteAmountValue() (int, error) {
	return configurationUseCase.configurationRepository.GetRealMaxVoteAmountValue()
}

func (configurationUseCase *ConfigurationUseCase) GetRealMaxConstancyAmountValue() (int, error) {
	return configurationUseCase.configurationRepository.GetRealMaxConstancyAmountValue()
}*/
