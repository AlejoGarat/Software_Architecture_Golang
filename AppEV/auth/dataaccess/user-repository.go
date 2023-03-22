package dataaccess

import (
	"auth/models/read"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	collection = "users"
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

func (userRepository UserRepository) FindUser(id string, password string) (read.User, error) {
	var user read.User
	query := bson.M{"id": id, "password": password}

	err := userRepository.mongoCli.Database(userRepository.db).Collection(collection).FindOne(context.TODO(), query).Decode(&user)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (userRepository UserRepository) AddTokenToUser(id string, token string) error {
	query := bson.M{"id": id}

	update := bson.M{"$set": bson.M{"token": token}}

	_, err := userRepository.mongoCli.Database(userRepository.db).Collection(collection).UpdateOne(context.TODO(), query, update)

	if err != nil {
		return err
	}

	return nil
}
