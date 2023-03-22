package dataaccess

import (
	"consultant-service/models/write"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type AlertRepository struct {
	redisCli *redis.Client
}

const (
	alertConfigKey        = "configElection"
	maxConstancyAmountKey = "maxConstancyAmount"
	mailRecipientsKey     = "mailRecipients"
)

func NewAlertRepo(redisCli *redis.Client) *AlertRepository {
	return &AlertRepository{
		redisCli: redisCli,
	}
}

func (alertRepository *AlertRepository) ModifyAlertConfiguration(alertConfig write.AlertConfiguration) error {
	jsonAlert, err := json.Marshal(alertConfig)

	if err != nil {
		return err
	}

	return alertRepository.redisCli.Set(alertRepository.redisCli.Context(), alertConfigKey+alertConfig.ElectionId, jsonAlert, 0).Err()
}

func (alertRepository *AlertRepository) GetAlertConfiguration(electionId string) (write.AlertConfiguration, error) {
	var alertConfig write.AlertConfiguration

	value, err := alertRepository.redisCli.Get(alertRepository.redisCli.Context(), alertConfigKey+electionId).Result()

	if err != nil {
		return alertConfig, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &alertConfig)
	}

	return alertConfig, err
}
