package dataaccess

import (
	"context"
	"crypto/sha256"
	"election-service/models/read"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type VoteRepository struct {
	mongoCli *mongo.Client
	db       string
}

func NewVoteMongoRepo(mongoCli *mongo.Client, db string) *VoteRepository {
	return &VoteRepository{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (voteRepository *VoteRepository) GetVoterVotingSchedules(electionId string, voterId string) ([]read.Info, error) {
	var vote read.Vote

	query := bson.M{"voter_id": hashData(voterId)}

	err := voteRepository.mongoCli.Database(voteRepository.db+electionId).Collection(votesCollection).FindOne(context.TODO(), query).Decode(&vote)

	if err != nil {
		return vote.Info, err
	}

	return vote.Info, nil
}

func hashData(data string) string {
	encryptedData := sha256.New()

	encryptedData.Write([]byte(data))

	return fmt.Sprintf("%x", encryptedData.Sum(nil))
}
