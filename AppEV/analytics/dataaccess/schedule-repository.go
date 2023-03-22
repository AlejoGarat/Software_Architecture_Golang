package dataaccess

import (
	"analytics/models/read"
	"context"
	"fmt"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	scheduleCollection = "voters_schedule"
)

type ScheduleRepository struct {
	mongoCli *mongo.Client
	db       string
}

func NewScheduleMongoRepo(mongoCli *mongo.Client, db string) *ScheduleRepository {
	return &ScheduleRepository{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (scheduleRepository *ScheduleRepository) GetFrequentSchedules(electionId string) (read.FrequentVotationSchedules, error) {
	var frequentVotationSchedules read.FrequentVotationSchedules

	query := bson.M{"election_id": electionId}

	err := scheduleRepository.mongoCli.Database(scheduleRepository.db).Collection(scheduleCollection).FindOne(context.TODO(), query).Decode(&frequentVotationSchedules)

	fmt.Println(err.Error())

	if err != nil {
		return frequentVotationSchedules, err
	}

	sort.SliceStable(frequentVotationSchedules.Schedules, func(i, j int) bool {
		return frequentVotationSchedules.Schedules[i].Votes > frequentVotationSchedules.Schedules[j].Votes
	})

	return frequentVotationSchedules, nil
}
