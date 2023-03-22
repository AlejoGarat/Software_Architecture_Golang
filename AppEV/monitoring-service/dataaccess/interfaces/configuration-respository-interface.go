package interfaces

import "monitoring-service/models"

type ConfigurationRepository interface {
	GetAlertConfiguration(string) (models.ConfigurationModel, error)
	GetRealMaxVoteAmountValue(string) (int, error)
	GetRealMaxConstancyAmountValue(string) (int, error)
	GetElections() ([]models.Election, error)
}
