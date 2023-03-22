package dataaccess

import (
	"context"
	"election-service/models/read"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userCollection = "users"
)

type UserRepository struct {
	mongoCli *mongo.Client
	db       string
}

func NewUserMongoRepo(mongoCli *mongo.Client, db string) *UserRepository {
	return &UserRepository{
		mongoCli: mongoCli,
		db:       db,
	}
}

func (userRepository *UserRepository) GetUserRole(id string, token string) (string, error) {
	var user read.User
	query := bson.M{"id": id, "token": token}

	err := userRepository.mongoCli.Database(userRepository.db).Collection(userCollection).FindOne(context.TODO(), query).Decode(&user)

	if err != nil {
		return user.Role, err
	}

	return user.Role, nil
}
