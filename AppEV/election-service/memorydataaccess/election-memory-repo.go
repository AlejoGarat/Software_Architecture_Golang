package memorydataaccess

import (
	"context"
	c "election-service/config"
	"election-service/models/read"
	"election-service/models/write"
	"encoding/json"
	"errors"
	"fmt"

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

func (electionRepo redisElectionImp) AddElection(election write.Election) error {

	jsonElection, _ := json.Marshal(election)

	err := electionRepo.redisCli.Set(context.TODO(), "electionId"+election.Id, string(jsonElection), 0).Err()

	return err
}

func (electionRepo redisElectionImp) AddVotersPerDepartment(votersPerDpt read.VotersPerDepartment) error {
	err := electionRepo.redisCli.Set(context.TODO(), "votersAmount"+votersPerDpt.ElectionId+votersPerDpt.Department, votersPerDpt.VotersAmount, 0).Err()

	return err
}

func (electionRepo redisElectionImp) AddCandidates(candidates []string, electionId string) error {
	jsonCandidates, err := json.Marshal(candidates)

	if err != nil {
		return err
	}

	err = electionRepo.redisCli.Set(context.TODO(), "candidates"+electionId, string(jsonCandidates), 0).Err()

	return err
}

func (electionRepo redisElectionImp) AddPoliticalParties(politicalParties []string, electionId string) error {
	jsonParties, err := json.Marshal(politicalParties)

	if err != nil {
		return err
	}

	err = electionRepo.redisCli.Set(context.TODO(), "politicalParties"+electionId, string(jsonParties), 0).Err()

	return err
}

func (electionRepo redisElectionImp) AddDepartments(departments []string, electionId string) error {
	jsonDepartments, err := json.Marshal(departments)

	if err != nil {
		return err
	}

	err = electionRepo.redisCli.Set(context.TODO(), "departments"+electionId, string(jsonDepartments), 0).Err()

	return err
}

func (electionRepo redisElectionImp) GetStartElectionFilters() (c.Configurations, error) {
	result := electionRepo.redisCli.Get(context.TODO(), "startElectionFilters")
	var filters c.Configurations

	if result.Err() == redis.Nil {
		return filters, errors.New("filters not in db")
	}

	err := json.Unmarshal([]byte(result.Val()), &filters)

	return filters, err
}

func (electionRepo redisElectionImp) GetCloseElectionFilters() (c.Configurations, error) {
	result := electionRepo.redisCli.Get(context.TODO(), "closeElectionFilters")
	var filters c.Configurations

	if result.Err() == redis.Nil {
		return filters, errors.New("filters not in db")
	}

	err := json.Unmarshal([]byte(result.Val()), &filters)

	return filters, err
}

func (electionRepo redisElectionImp) GetElectionResult(electionId string) (read.ElectionResult, error) {
	var electionResult read.ElectionResult

	electionVoting, err := electionRepo.getElectionVotingData(electionId)

	if err != nil {
		return electionResult, err
	}
	fmt.Println("1")

	electionResult.ElectionVotingData = electionVoting

	candidatesVotes, err := electionRepo.getVotesPerCandidate(electionId)

	if err != nil {
		return electionResult, err
	}

	electionResult.VotesPerCandidate = candidatesVotes

	partyVotes, err := electionRepo.getVotesPerPoliticalParty(electionId)

	if err != nil {

		return electionResult, err
	}
	fmt.Println("3")
	electionResult.VotesPerPoliticalParty = partyVotes

	departmentEligibleVoters, departmentVotes, err := electionRepo.getDepartmentVoting(electionId)

	if err != nil {

		return electionResult, err
	}
	fmt.Println("4")
	electionResult.EligibleVotersPerDepartment = departmentEligibleVoters
	electionResult.VotesPerDepartment = departmentVotes

	return electionResult, nil
}

func (electionRepo redisElectionImp) getElectionVotingData(electionId string) (read.ElectionVotingData, error) {
	result := electionRepo.redisCli.Get(context.TODO(), "electionId"+electionId)
	var electionVoting read.ElectionVotingData

	if result.Err() == redis.Nil {
		return electionVoting, errors.New("error getting election")
	}

	err := json.Unmarshal([]byte(result.Val()), &electionVoting)

	if err != nil {
		return electionVoting, err
	}

	return electionVoting, nil
}

func (electionRepo redisElectionImp) getVotesPerCandidate(electionId string) ([]read.StringIntCandidate, error) {
	result := electionRepo.redisCli.Get(context.TODO(), "votesPerCandidates"+electionId)
	var candidatesVotes []read.StringIntCandidate

	if result.Err() == redis.Nil {
		return candidatesVotes, errors.New("error getting candidates")
	}

	err := json.Unmarshal([]byte(result.Val()), &candidatesVotes)

	return candidatesVotes, err
}

func (electionRepo redisElectionImp) getVotesPerPoliticalParty(electionId string) ([]read.StringIntParty, error) {
	result := electionRepo.redisCli.Get(context.TODO(), "votesPerPoliticalParties"+electionId)
	var partiesVotes []read.StringIntParty

	if result.Err() == redis.Nil {
		return partiesVotes, errors.New("error getting parties")
	}

	err := json.Unmarshal([]byte(result.Val()), &partiesVotes)

	return partiesVotes, err
}

func (electionRepo redisElectionImp) getDepartmentVoting(electionId string) ([]read.StringIntEligibleDepartment, []read.StringIntDepartment, error) {
	result := electionRepo.redisCli.Get(context.TODO(), "departments"+electionId)

	var departmentsEligibleVoters []read.StringIntEligibleDepartment
	var departmentsVotes []read.StringIntDepartment

	var departments []string

	if result.Err() == redis.Nil {
		return departmentsEligibleVoters, departmentsVotes, errors.New("error getting departments")
	}

	err := json.Unmarshal([]byte(result.Val()), &departments)

	if err != nil {
		return departmentsEligibleVoters, departmentsVotes, err
	}

	for _, department := range departments {
		result := electionRepo.redisCli.Get(context.TODO(), "votersAmount"+electionId+department)
		var amount int

		if result.Err() == redis.Nil {
			return departmentsEligibleVoters, departmentsVotes, errors.New("error getting voters of department")
		}

		err = json.Unmarshal([]byte(result.Val()), &amount)

		if err != nil {
			return departmentsEligibleVoters, departmentsVotes, errors.New("error getting voters of department")
		}

		departmentEligibleVoters := read.StringIntEligibleDepartment{Id: department, Amount: amount}

		departmentsEligibleVoters = append(departmentsEligibleVoters, departmentEligibleVoters)
	}

	result = electionRepo.redisCli.Get(context.TODO(), "votesPerDepartments"+electionId)

	if result.Err() == redis.Nil {
		return departmentsEligibleVoters, departmentsVotes, errors.New("error getting departments")
	}

	err = json.Unmarshal([]byte(result.Val()), &departmentsVotes)

	return departmentsEligibleVoters, departmentsVotes, err
}
