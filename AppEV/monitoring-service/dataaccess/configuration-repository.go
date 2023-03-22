package dataaccess

import (
	"context"
	"encoding/json"
	"monitoring-service/models"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	alertConfigKey      = "configElection"
	votesCollection     = "votes"
	constancyCollection = "constancy"
	electionsCollection = "election"
	electionDB          = "ElectionDB_"
)

type ConfigurationRepository struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
	db       string
}

func NewConfigurationRepository(mongoCli *mongo.Client, redisCli *redis.Client, db string) *ConfigurationRepository {
	return &ConfigurationRepository{
		mongoCli: mongoCli,
		redisCli: redisCli,
		db:       db,
	}
}

func (alertRepository *ConfigurationRepository) GetAlertConfiguration(electionId string) (models.ConfigurationModel, error) {
	var alertConfig models.ConfigurationModel

	value, err := alertRepository.redisCli.Get(alertRepository.redisCli.Context(), alertConfigKey+electionId).Result()

	if err != nil {
		return models.ConfigurationModel{}, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &alertConfig)
	}

	return alertConfig, err
}

func (alertRepository *ConfigurationRepository) GetRealMaxVoteAmountValue(electionId string) (int, error) {

	cursorVotes, err := alertRepository.mongoCli.Database(electionDB+electionId).Collection(votesCollection).Find(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	defer cursorVotes.Close(context.TODO())

	var maxVoteAmount int
	maxVoteAmount = 0

	for cursorVotes.Next(context.TODO()) {
		var vote models.Vote
		var voteInfo []models.Info
		var voteInfoSize int

		if err = cursorVotes.Decode(&vote); err != nil {
			return 0, err
		}

		voteInfo = vote.CompleteInfo
		voteInfoSize = len(voteInfo)

		if voteInfoSize > maxVoteAmount {
			maxVoteAmount = voteInfoSize
		}
	}

	return maxVoteAmount, nil
}

func (alertRepository *ConfigurationRepository) GetRealMaxConstancyAmountValue(electionId string) (int, error) {

	cursorConstancy, err := alertRepository.mongoCli.Database(electionDB+electionId).Collection(constancyCollection).Find(context.TODO(), bson.D{})

	if err != nil {
		return 0, err
	}

	defer cursorConstancy.Close(context.TODO())

	var maxConstancyAmount int
	maxConstancyAmount = 0

	for cursorConstancy.Next(context.TODO()) {
		var constancy models.Constancy
		var constancyTimeStampsSize int

		if err = cursorConstancy.Decode(&constancy); err != nil {
			return 0, err
		}

		constancyTimeStampsSize = len(constancy.Timestamps)

		if constancyTimeStampsSize > maxConstancyAmount {
			maxConstancyAmount = constancyTimeStampsSize
		}
	}

	return maxConstancyAmount, nil
}

func (alertRepository *ConfigurationRepository) GetElections() ([]models.Election, error) {
	var elections []models.Election

	cursorElections, err := alertRepository.mongoCli.Database(alertRepository.db).Collection(electionsCollection).Find(context.TODO(), bson.D{})

	if err != nil {
		return elections, err
	}

	err = cursorElections.All(context.TODO(), &elections)

	return elections, err
}
