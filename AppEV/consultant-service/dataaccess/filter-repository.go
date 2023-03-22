package dataaccess

import (
	"consultant-service/models/read"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

const (
	electionBeginningFiltersKey = "startElectionFilters"
	electionEndFiltersKey       = "closeElectionFilters"
	voteIssuanceFiltersKey      = "voteFilters"
)

type FilterRepository struct {
	redisCli *redis.Client
}

func NewFilterRepo(redisCli *redis.Client) *FilterRepository {
	return &FilterRepository{
		redisCli: redisCli,
	}
}

func (filterRepository *FilterRepository) ModifyElectionBeginningFilters(filters read.ElectionBeginningFilters) error {
	jsonFilters, err := json.Marshal(filters)

	if err != nil {
		return err
	}

	err = filterRepository.redisCli.Set(filterRepository.redisCli.Context(), electionBeginningFiltersKey, jsonFilters, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (filterRepository *FilterRepository) ModifyElectionEndFilters(filters read.ElectionEndFilters) error {
	jsonFilters, err := json.Marshal(filters)

	if err != nil {
		return err
	}

	err = filterRepository.redisCli.Set(filterRepository.redisCli.Context(), electionEndFiltersKey, jsonFilters, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (filterRepository *FilterRepository) ModifyVoteIssuanceFilters(filters read.VoteIssuanceFilters) error {
	jsonFilters, err := json.Marshal(filters)

	if err != nil {
		return err
	}

	err = filterRepository.redisCli.Set(filterRepository.redisCli.Context(), voteIssuanceFiltersKey, jsonFilters, 0).Err()

	if err != nil {
		return err
	}

	return nil
}
