package dataaccess

import (
	"encoding/json"
	"votation-service/models/read"

	"github.com/go-redis/redis/v8"
)

const (
	filtersKey = "voteFilters"
)

type FilterRepository struct {
	redisCli *redis.Client
}

func NewFilterRepository(redisCli *redis.Client) *FilterRepository {
	return &FilterRepository{
		redisCli: redisCli,
	}
}

func (filterRepository *FilterRepository) GetFilters() (read.VoteFilters, error) {
	var filters read.VoteFilters

	value, err := filterRepository.redisCli.Get(filterRepository.redisCli.Context(), filtersKey).Result()

	if err != nil {
		return filters, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &filters)
	}

	return filters, err
}
