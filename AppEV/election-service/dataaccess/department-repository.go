package dataaccess

import (
	"context"
	"election-service/models/read"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	departmentCollection = "voters_per_department"
)

type DepartmentRepository struct {
	mongoCli *mongo.Client
	db       string
}

func NewDepartmentMongoRepo(mongoCli *mongo.Client, db string) *DepartmentRepository {
	return &DepartmentRepository{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (departmentRepository *DepartmentRepository) GetVoteCoveragePerDepartment(electionId string) ([]read.DepartmentVoteCoverage, error) {
	var voteCoverage []read.DepartmentVoteCoverage

	query := bson.M{"election_id": electionId}

	cursor, err := departmentRepository.mongoCli.Database(departmentRepository.db).Collection(departmentCollection).Find(context.TODO(), query)

	if err != nil {
		return voteCoverage, err
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var voteCoverageDepartment read.DepartmentVoteCoverage

		if err = cursor.Decode(&voteCoverageDepartment); err != nil {
			return voteCoverage, nil
		}

		voteCoverage = append(voteCoverage, voteCoverageDepartment)
	}

	return voteCoverage, nil
}
