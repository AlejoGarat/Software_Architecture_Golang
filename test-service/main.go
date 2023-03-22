package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"test-service/datasources"
	"test-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	iterations      = 100
	votationService = "http://127.0.0.1:8084/votation-api/v1/vote"
)

const (
	mongoAddress = "mongodb://localhost:27017/"
	redisAddress = "localhost:6379"
)

func main() {

	mongoClient, err := datasources.NewMongoDataSource(mongoAddress)

	if err != nil {
		panic(err)
	}

	defer mongoClient.Disconnect(context.TODO())

	Start(mongoClient)
}

func Start(mongoCli *mongo.Client) {
	election := GetCompleteElection(mongoCli)

	RunJobs(election)
}

func RunJobs(election models.CompleteElection) {

	for i := 0; i <= 50; i++ {
		go func() {
			for j := 0; j <= iterations/50; j++ {
				for _, voter := range election.Voters {
					data := map[string]string{
						"election_id":  "1",
						"voter_id":     voter.Id,
						"circuit_id":   voter.CircuitId,
						"candidate_id": "96561059",
					}

					jsonData, _ := json.Marshal(data)

					resp, _ := http.Post(votationService, "application/json", bytes.NewBuffer(jsonData))

					body, _ := ioutil.ReadAll(resp.Body)

					fmt.Println(string(body))
				}
			}
		}()
	}

	for {
	}
}

func GetCompleteElection(mongoCli *mongo.Client) models.CompleteElection {
	var election models.CompleteElection

	voters, err := GetVotes(mongoCli)

	if err != nil {
		panic(err)
	}

	election.Voters = voters

	return election
}

func GetVotes(mongoCli *mongo.Client) ([]models.Voter, error) {
	var voters []models.Voter

	cursorVoters, err := mongoCli.Database("ElectionDB_1").Collection("voters").Find(context.TODO(), bson.M{})

	if err != nil {
		return voters, err
	}

	err = cursorVoters.All(context.TODO(), &voters)

	return voters, err
}
