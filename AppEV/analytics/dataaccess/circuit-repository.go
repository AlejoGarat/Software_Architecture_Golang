package dataaccess

import (
	"analytics/models/read"

	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	circuitCollection    = "circuits"
	ageVoteCollection    = "voters_per_age"
	genderVoteCollection = "voters_per_gender"
)

type CircuitRepository struct {
	mongoCli *mongo.Client
	db       string
}

func NewCircuitMongoRepo(mongoCli *mongo.Client, db string) *CircuitRepository {
	return &CircuitRepository{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (circuitRepository *CircuitRepository) GetVoteCoveragePerCircuit(electionId string) ([]read.CircuitVoteCoverage, error) {
	var voteCoverage []read.CircuitVoteCoverage

	cursor, err := circuitRepository.mongoCli.Database(circuitRepository.db).Collection(circuitCollection).Find(context.TODO(), bson.M{})

	if err != nil {
		return voteCoverage, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var circuitVoteCoverage read.CircuitVoteCoverage
		var circuit read.Circuit

		if err = cursor.Decode(&circuit); err != nil {
			return voteCoverage, err
		}

		circuitVoteCoverage.CircuitId = circuit.CircuitId

		query := bson.M{"circuit_id": circuit.CircuitId}

		err := circuitRepository.mongoCli.Database(circuitRepository.db).Collection(ageVoteCollection).FindOne(context.TODO(), query).Decode(&circuitVoteCoverage.AgeVote)

		if err != nil {
			return voteCoverage, err
		}

		err = circuitRepository.mongoCli.Database(circuitRepository.db).Collection(genderVoteCollection).FindOne(context.TODO(), query).Decode(&circuitVoteCoverage.GenderVote)

		if err != nil {
			return voteCoverage, err
		}

		voteCoverage = append(voteCoverage, circuitVoteCoverage)
	}

	return voteCoverage, nil
}
