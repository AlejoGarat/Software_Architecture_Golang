package dataaccess

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"votation-service/models/read"
	"votation-service/models/write"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	electionDB          = "ElectionDB_"
	votesCollection     = "votes"
	voterCollection     = "voters"
	candidateCollection = "candidates"
	electionCollection  = "election"
	constancyCollection = "constancy"
	redisElectionKey    = "electionId"
)

type VoteRepository struct {
	mongoCli *mongo.Client
	redisCli *redis.Client
	db       string
}

func NewVotesRepo(mongoCli *mongo.Client, redisCli *redis.Client, db string) *VoteRepository {
	return &VoteRepository{
		mongoCli: mongoCli,
		redisCli: redisCli,
		db:       db,
	}
}

func (voteDb *VoteRepository) AddVote(vote write.Vote) error {
	vote = hashVote(vote)
	vote.EndingTime = time.Now()
	vote.ProcessingTime = fmt.Sprintf("%f", vote.EndingTime.Sub(vote.StartingTime).Seconds()) + " s"

	query := bson.M{"voter_id": vote.VoterDocument}
	update := bson.M{"$set": bson.M{"id": vote.VoteId, "candidate_id": vote.CandidateId, "circuit_id": vote.CircuitId, "starting_time": vote.StartingTime,
		"ending_time": vote.EndingTime, "processing_time": vote.ProcessingTime}, "$push": bson.M{"info": bson.M{"date": vote.Info.Date, "hour": vote.Info.Hour}}}
	opts := options.Update().SetUpsert(true)

	_, err := voteDb.mongoCli.Database(electionDB+vote.ElectionId).Collection(votesCollection).UpdateOne(context.TODO(), query, update, opts)

	if err != nil {
		return err
	}

	return nil
}

func (voteDb *VoteRepository) GetVoterById(document string, electionDBId string) (write.Voter, error) {
	var voter write.Voter
	query := bson.M{"id": document}

	err := voteDb.mongoCli.Database(electionDB+electionDBId).Collection(voterCollection).FindOne(context.TODO(), query).Decode(&voter)

	return voter, err
}

func (voteDb *VoteRepository) GetCandidateById(id string, electionDBId string) (write.Candidate, error) {
	var candidate write.Candidate
	query := bson.M{"id": id}

	err := voteDb.mongoCli.Database(electionDB+electionDBId).Collection(candidateCollection).FindOne(context.TODO(), query).Decode(&candidate)

	return candidate, err
}

func (voteDb *VoteRepository) GetElectionById(id string) (read.Election, error) {
	var election read.Election

	value, err := voteDb.redisCli.Get(voteDb.redisCli.Context(), redisElectionKey+id).Result()

	if err != nil {
		return read.Election{}, err
	}

	if err != redis.Nil {
		err = json.Unmarshal([]byte(value), &election)
	}

	return election, err
}

func (voteDb *VoteRepository) GetVoteById(id string, electionId string) (write.Vote, error) {
	var vote write.Vote
	query := bson.M{"id": id}

	err := voteDb.mongoCli.Database(electionDB+electionId).Collection(votesCollection).FindOne(context.TODO(), query).Decode(&vote)

	return vote, err
}

func (voteDb *VoteRepository) ExistsVoteOfVoter(document string, electionId string) bool {
	query := bson.M{"voter_id": hashData(document)}

	err := voteDb.mongoCli.Database(electionDB+electionId).Collection(votesCollection).FindOne(context.TODO(), query).Err()

	return err == nil
}

func (voteDb *VoteRepository) SetConstancyData(constancyDBData write.ConstancyDBData, electionId string) error {
	query := bson.M{"voter_id": constancyDBData.VoterDocument}
	update := bson.M{"$push": bson.M{"timestamps": constancyDBData.Timestamp}}
	opts := options.Update().SetUpsert(true)

	_, err := voteDb.mongoCli.Database(electionDB+electionId).Collection(constancyCollection).UpdateByID(context.TODO(), query, update, opts)

	if err != nil {
		return err
	}

	return err
}

func (voteDb *VoteRepository) GetConstancyData(constancyReq read.ConstancyRequest) (write.Constancy, error) {
	election, err := voteDb.GetElectionById(constancyReq.ElectionId)

	if err != nil {
		return write.Constancy{}, err
	}

	voterDocument := constancyReq.VoterId

	voter, err := voteDb.GetVoterById(constancyReq.VoterId, constancyReq.ElectionId)

	if err != nil {
		return write.Constancy{}, err
	}

	vote, err := voteDb.GetVoteById(constancyReq.VoteId, constancyReq.ElectionId)

	if err != nil {
		return write.Constancy{}, err
	}

	if vote.VoterDocument != hashData(voterDocument) {
		return write.Constancy{}, errors.New("The vote does not correspond to the voter.")
	}

	constancy := write.Constancy{
		Election:             election,
		VoteEmissionInfoDate: vote.InfoArr[len(vote.InfoArr)-1],
		VoterDocument:        voterDocument,
		Name:                 voter.Name,
		Surname:              voter.Surname,
		VoteId:               vote.VoteId,
	}

	return constancy, nil
}

func hashData(data string) string {
	encryptedData := sha256.New()

	encryptedData.Write([]byte(data))

	return fmt.Sprintf("%x", encryptedData.Sum(nil))
}

func hashVote(vote write.Vote) write.Vote {
	vote.CandidateId = hashData(vote.CandidateId)
	vote.VoterDocument = hashData(vote.VoterDocument)

	return vote
}
