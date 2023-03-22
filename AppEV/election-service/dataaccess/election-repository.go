package dataaccess

import (
	"context"
	"election-service/models/read"
	"election-service/models/write"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	electoralDbName               = "ElectoralDB"
	electionDbName                = "ElectionDB_"
	electionCollection            = "election"
	votersCollection              = "voters"
	votesCollection               = "votes"
	candidatesCollection          = "candidates"
	politicalPartiesCollection    = "political_party"
	circuitsCollection            = "circuit"
	votersPerAgeCollection        = "voters_per_age"
	votersPerGenderCollection     = "voters_per_gender"
	votersPerDepartmentCollection = "voters_per_department"
)

type electionRepoMongoImp struct {
	mongoCli *mongo.Client
	db       string
}

func NewElectionMongoRepo(mongoCli *mongo.Client, db string) *electionRepoMongoImp {
	return &electionRepoMongoImp{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (electionDb *electionRepoMongoImp) AddElection(election write.Election) error {
	_, err := electionDb.mongoCli.Database(electoralDbName).Collection(electionCollection).InsertOne(context.TODO(), election)
	return err
}

func (electionDb *electionRepoMongoImp) AddVoter(voter read.Voter, electionId string) error {

	_, err := electionDb.mongoCli.Database(electionDbName+electionId).Collection(votersCollection).InsertOne(context.TODO(), voter)

	return err
}

func (electionDb *electionRepoMongoImp) AddCandiadtes(candidates []write.Candidate, electionId string) error {
	for _, candidate := range candidates {
		_, err := electionDb.mongoCli.Database(electoralDbName).Collection(candidatesCollection).InsertOne(context.TODO(), candidate)

		if err != nil {
			return err
		}

		read_candidate := read.Candidate{
			CandidateId:      candidate.CandidateId,
			PoliticalPartyId: candidate.PoliticalPartyId,
		}

		_, err = electionDb.mongoCli.Database(electionDbName+electionId).Collection(candidatesCollection).InsertOne(context.TODO(), read_candidate)

		if err != nil {
			return err
		}
	}

	return nil
}

func (electionDb *electionRepoMongoImp) AddPoliticalParties(political_parties []write.PoliticalParty, electionId string) error {
	for _, political_party := range political_parties {
		political_party.ElectionId = electionId
		_, err := electionDb.mongoCli.Database(electoralDbName).Collection(politicalPartiesCollection).InsertOne(context.TODO(), political_party)

		if err != nil {
			return err
		}
	}

	return nil
}

func (electionDb *electionRepoMongoImp) AddCircuits(circuits []write.Circuit) error {
	for _, circuit := range circuits {
		_, err := electionDb.mongoCli.Database(electoralDbName).Collection(circuitsCollection).InsertOne(context.TODO(), circuit)

		if err != nil {
			return err
		}
	}

	return nil
}

func (electionDb *electionRepoMongoImp) GetVotes(electionId string) ([]read.Vote, error) {
	var votes []read.Vote

	cursorVotes, err := electionDb.mongoCli.Database(electionDbName+electionId).Collection(votesCollection).Find(context.TODO(), bson.M{})

	if err != nil {
		return votes, err
	}

	err = cursorVotes.All(context.TODO(), &votes)

	return votes, err
}

func (electionDb *electionRepoMongoImp) GetCandidate(id string) (read.Candidate, error) {
	var candidate read.Candidate
	query := bson.M{"candidate_id": id}

	err := electionDb.mongoCli.Database(electoralDbName).Collection(candidatesCollection).FindOne(context.TODO(), query).Decode(&candidate)

	return candidate, err
}

func (electionDb *electionRepoMongoImp) AddVotersPerAge(votersPerAge map[string]map[int]int, electionId string) error {

	for circuit, ageVoters := range votersPerAge {

		votersPerAgeValue := read.VotersPerAge{
			ElectionId: electionId,
			CircuitId:  circuit,
			Voters:     []read.AgeAmount{},
		}

		for age, votersAmount := range ageVoters {
			votersPerAgeValue.Voters = append(votersPerAgeValue.Voters, read.AgeAmount{
				Age:          age,
				VotersAmount: votersAmount,
			})
		}

		_, err := electionDb.mongoCli.Database(electoralDbName).Collection(votersPerAgeCollection).InsertOne(context.TODO(), votersPerAgeValue)

		if err != nil {
			return err
		}
	}

	return nil
}

func (electionDb *electionRepoMongoImp) AddVotersPerGender(votersPerGender map[string]map[string]int, electionId string) error {

	for circuit, genderVoters := range votersPerGender {

		votersPerGenderValue := read.VotersPerGender{
			ElectionId: electionId,
			CircuitId:  circuit,
			Voters:     []read.GenderAmount{},
		}

		for gender, votersAmount := range genderVoters {
			votersPerGenderValue.Voters = append(votersPerGenderValue.Voters, read.GenderAmount{
				Gender:       gender,
				VotersAmount: votersAmount,
			})
		}

		_, err := electionDb.mongoCli.Database(electoralDbName).Collection(votersPerGenderCollection).InsertOne(context.TODO(), votersPerGenderValue)

		if err != nil {
			return err
		}
	}

	return nil
}

func (electionDb *electionRepoMongoImp) AddVotersPerDepartmentByGenderAndAge(votersPerGender map[string]map[string]int, votersPerAge map[string]map[int]int, electionId string) error {

	for department, genderVoters := range votersPerGender {

		votersPerDepartment := read.VotersPerDepartmentByGenderAndAge{
			ElectionId:      electionId,
			Department:      department,
			VotersPerAge:    []read.AgeAmount{},
			VotersPerGender: []read.GenderAmount{},
		}

		for gender, votersAmount := range genderVoters {
			votersPerDepartment.VotersPerGender = append(votersPerDepartment.VotersPerGender, read.GenderAmount{
				Gender:       gender,
				VotersAmount: votersAmount,
			})
		}

		for age, votersAmount := range votersPerAge[department] {
			votersPerDepartment.VotersPerAge = append(votersPerDepartment.VotersPerAge, read.AgeAmount{
				Age:          age,
				VotersAmount: votersAmount,
			})
		}

		_, err := electionDb.mongoCli.Database(electoralDbName).Collection(votersPerDepartmentCollection).InsertOne(context.TODO(), votersPerDepartment)

		if err != nil {
			return err
		}
	}

	return nil
}
