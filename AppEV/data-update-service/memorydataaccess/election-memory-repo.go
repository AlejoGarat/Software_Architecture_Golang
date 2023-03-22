package memorydataaccess

import (
	"context"
	"data-update-service/models/read"
	"encoding/json"
	"errors"

	redis "github.com/go-redis/redis/v8"
)

type redisElectionImp struct {
	redisCli *redis.Client
}

func NewRedisElectionImp(redisCli *redis.Client) *redisElectionImp {
	return &redisElectionImp{
		redisCli: redisCli,
	}
}

func (mr redisElectionImp) UpdateVotesPerParty(votesPerParty []read.StringIntParty, electionId string) error {
	jsonVotesPerParty, _ := json.Marshal(votesPerParty)

	err := mr.redisCli.Set(context.TODO(), "votesPerPoliticalParties"+electionId, string(jsonVotesPerParty), 0).Err()

	return err
}

func (mr redisElectionImp) UpdateVotesPerCandidate(votesPerCandidate []read.StringIntCandidate, electionId string) error {
	jsonVotesPerCandidate, _ := json.Marshal(votesPerCandidate)

	err := mr.redisCli.Set(context.TODO(), "votesPerCandidates"+electionId, string(jsonVotesPerCandidate), 0).Err()

	return err
}

func (mr redisElectionImp) UpdateVotesPerDepartment(votesPerDepartment []read.StringIntDepartment, electionId string) error {
	jsonVotesPerDepartment, _ := json.Marshal(votesPerDepartment)

	err := mr.redisCli.Set(context.TODO(), "votesPerDepartments"+electionId, string(jsonVotesPerDepartment), 0).Err()

	return err
}

func (mr redisElectionImp) UpdateTotalVotes(electionId string) error {
	var completeElection read.CompleteElection
	result := mr.redisCli.Get(context.TODO(), "electionId"+electionId)

	if result.Err() == redis.Nil {
		return errors.New("error getting election")
	}

	err := json.Unmarshal([]byte(result.Val()), &completeElection)

	if err != nil {
		return err
	}

	completeElection.Votes = completeElection.Votes + 1

	jsonElection, err := json.Marshal(completeElection)

	if err != nil {
		return err
	}

	err = mr.redisCli.Set(context.TODO(), "electionId"+electionId, string(jsonElection), 0).Err()

	return err
}
