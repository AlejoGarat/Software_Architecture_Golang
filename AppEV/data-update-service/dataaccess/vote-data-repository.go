package dataaccess

import (
	"context"
	"data-update-service/models/read"
	"data-update-service/models/write"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	electionDB           = "ElectionDB_"
	electionCollection   = "election"
	scheduleCollection   = "voters_schedule"
	ageVoteCollection    = "voters_per_age"
	genderVoteCollection = "voters_per_gender"
	departmentCollection = "voters_per_department"
	votersCollection     = "voters"
	votesCollection      = "votes"
	candidatesCollection = "candidates"
	circuitCollection    = "circuit"
)

type VoteDataRepository struct {
	mongoCli *mongo.Client
	db       string
}

func NewVoteDataRepository(mongoCli *mongo.Client, db string) *VoteDataRepository {
	return &VoteDataRepository{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (voteDataRepository *VoteDataRepository) UpdateSchedule(vote read.Vote) error {
	var schedule write.FrequentVotationSchedules

	query := bson.M{"election_id": vote.ElectionId}

	err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(scheduleCollection).FindOne(context.TODO(), query).Decode(&schedule)

	if err != nil {
		hourVote := write.HourVote{Hour: vote.Info.Hour, Votes: 1}

		schedule = write.FrequentVotationSchedules{
			ElectionId: vote.ElectionId,
			Schedules:  []write.HourVote{hourVote},
		}

		_, err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(scheduleCollection).InsertOne(context.TODO(), schedule)

		return err
	}

	voteInfoExists := false

	for i, voteInfo := range schedule.Schedules {
		if voteInfo.Hour == vote.Info.Hour {
			schedule.Schedules[i].Votes = voteInfo.Votes + 1
			voteInfoExists = true
			break
		}
	}

	var update primitive.M

	if voteInfoExists {
		update = bson.M{"$set": bson.M{"info": schedule.Schedules}}
	} else {
		info := write.HourVote{
			Hour:  vote.Info.Hour,
			Votes: 1,
		}
		update = bson.M{"$push": bson.M{"info": info}}
	}

	_, err = voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(scheduleCollection).UpdateOne(context.TODO(), query, update)

	return err
}

func (voteDataRepository *VoteDataRepository) UpdateCircuitAgeVotes(vote read.Vote, voter read.Voter) error {
	var ageVotes write.AgeVote

	query := bson.M{"election_id": vote.ElectionId, "circuit_id": vote.CircuitId}

	err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(ageVoteCollection).FindOne(context.TODO(), query).Decode(&ageVotes)

	if err != nil {
		return err
	}

	voteAgeExists := false

	for i, ageInfo := range ageVotes.CircuitVotesPerAge {
		if ageInfo.Age == voter.Age {
			ageVotes.CircuitVotesPerAge[i].Votes = ageInfo.Votes + 1
			voteAgeExists = true
			break
		}
	}

	var update primitive.M

	if voteAgeExists {
		update = bson.M{"$set": bson.M{"votes_age": ageVotes.CircuitVotesPerAge}}

	} else {
		age := write.VotesPerAge{
			Age:   voter.Age,
			Votes: 1,
		}
		update = bson.M{"$push": bson.M{"votes_age": age}}
	}

	_, err = voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(ageVoteCollection).UpdateOne(context.TODO(), query, update)

	return err
}

func (voteDataRepository *VoteDataRepository) UpdateCircuitGenderVotes(vote read.Vote, voter read.Voter) error {
	var genderVotes write.GenderVote

	query := bson.M{"election_id": vote.ElectionId, "circuit_id": vote.CircuitId}

	err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(genderVoteCollection).FindOne(context.TODO(), query).Decode(&genderVotes)

	if err != nil {
		return err
	}

	voteGenderExists := false

	for i, genderInfo := range genderVotes.CircuitVotesPerGender {
		if genderInfo.Gender == voter.Sex {
			genderVotes.CircuitVotesPerGender[i].Votes = genderInfo.Votes + 1
			voteGenderExists = true
			break
		}
	}

	var update primitive.M

	if voteGenderExists {
		update = bson.M{"$set": bson.M{"votes_gender": genderVotes.CircuitVotesPerGender}}
	} else {
		gender := write.VotesPerGender{
			Gender: voter.Sex,
			Votes:  1,
		}
		update = bson.M{"$push": bson.M{"votes_gender": gender}}
	}

	_, err = voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(genderVoteCollection).UpdateOne(context.TODO(), query, update)

	return err
}

func (voteDataRepository *VoteDataRepository) UpdateDepartmentData(vote read.Vote, voter read.Voter) error {
	var departmentCoverage write.DepartmentVoteCoverage

	query := bson.M{"election_id": vote.ElectionId, "department": voter.Department}

	err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(departmentCollection).FindOne(context.TODO(), query).Decode(&departmentCoverage)

	if err != nil {
		return err
	}

	voterGenderExists := false
	for i, genderInfo := range departmentCoverage.DepartmentVotesPerGender {
		if genderInfo.Gender == voter.Sex {
			departmentCoverage.DepartmentVotesPerGender[i].Votes = genderInfo.Votes + 1
			voterGenderExists = true
			break
		}
	}
	var update primitive.M

	if voterGenderExists {
		update = bson.M{"$set": bson.M{"votes_gender": departmentCoverage.DepartmentVotesPerGender}}
	} else {
		votesGender := write.VotesPerGender{
			Gender: voter.Sex,
			Votes:  1,
		}
		update = bson.M{"$push": bson.M{"votes_gender": votesGender}}
	}

	_, err = voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(departmentCollection).UpdateOne(context.TODO(), query, update)

	if err != nil {
		return err
	}

	voterAgeExists := false

	for i, ageInfo := range departmentCoverage.DepartmentVotesPerAge {
		if ageInfo.Age == voter.Age {
			departmentCoverage.DepartmentVotesPerAge[i].Votes = ageInfo.Votes + 1
			voterAgeExists = true
			break
		}
	}

	if voterAgeExists {
		update = bson.M{"$set": bson.M{"votes_age": departmentCoverage.DepartmentVotesPerAge}}
	} else {
		votesAge := write.VotesPerAge{
			Age:   voter.Age,
			Votes: 1,
		}
		update = bson.M{"$push": bson.M{"votes_age": votesAge}}
	}

	_, err = voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(departmentCollection).UpdateOne(context.TODO(), query, update)

	return err
}

func (voteDataRepository *VoteDataRepository) VoterHasVoted(voterDocument string, electionDbId string) bool {
	var schedule write.FrequentVotationSchedules
	query := bson.M{"voter_id": voterDocument}

	err := voteDataRepository.mongoCli.Database(electionDB+electionDbId).Collection(votesCollection).FindOne(context.TODO(), query).Decode(&schedule)

	if err != nil {
		return true
	}

	voteInfoLength := len(schedule.Schedules)

	return voteInfoLength > 1
}

func (voteDataRepository *VoteDataRepository) GetVoterByDocument(voterDocument string, electionDbId string) (read.Voter, error) {
	var voter read.Voter
	query := bson.M{"id": voterDocument}

	err := voteDataRepository.mongoCli.Database(electionDB+electionDbId).Collection(votersCollection).FindOne(context.TODO(), query).Decode(&voter)

	return voter, err
}

func (voteDataRepository *VoteDataRepository) GetVotes(electionDbId string) ([]read.VoteGet, error) {
	var votes []read.VoteGet

	cursorVotes, err := voteDataRepository.mongoCli.Database(electionDB+electionDbId).Collection(votesCollection).Find(context.TODO(), bson.M{})

	if err != nil {
		return votes, err
	}

	err = cursorVotes.All(context.TODO(), &votes)

	return votes, err
}

func (voteDataRepository *VoteDataRepository) GetCircuits(electionId string) ([]read.Circuit, error) {
	var circuit []read.Circuit

	cursorCircuits, err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(circuitCollection).Find(context.TODO(), bson.M{})

	if err != nil {
		return circuit, err
	}

	err = cursorCircuits.All(context.TODO(), &circuit)

	return circuit, err
}

func (voteDataRepository *VoteDataRepository) GetCandidates(electionId string) ([]read.Candidate, error) {
	var candidates []read.Candidate

	cursorCandidate, err := voteDataRepository.mongoCli.Database(electionDB+electionId).Collection(candidatesCollection).Find(context.TODO(), bson.M{})

	if err != nil {
		return candidates, err
	}

	err = cursorCandidate.All(context.TODO(), &candidates)

	return candidates, err
}

func (voteDataRepository *VoteDataRepository) UpdateTotalVotes(electionId string) error {
	query := bson.M{"id": electionId}
	update := bson.M{"$inc": bson.M{"total_votes": 1}}

	_, err := voteDataRepository.mongoCli.Database(voteDataRepository.db).Collection(electionCollection).UpdateOne(context.TODO(), query, update)

	return err
}
