package dataaccess

import (
	"election-service/models/read"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type AlertRepository struct {
	redisCli *redis.Client
}

const (
	alertConfigKey = "configElection"
)

func NewAlertRepo(redisCli *redis.Client) *AlertRepository {
	return &AlertRepository{
		redisCli: redisCli,
	}
}

func (alertRepository *AlertRepository) GetAlertConfiguration(electionId string) (read.AlertConfiguration, error) {
	var alertConfig read.AlertConfiguration

	value, err := alertRepository.redisCli.Get(alertRepository.redisCli.Context(), alertConfigKey+electionId).Result()

	if err != nil {
		return alertConfig, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &alertConfig)
	}

	return alertConfig, err
}
