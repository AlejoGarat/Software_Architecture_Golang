package factory

import (
	"auth/api/controllers"
	"auth/dataaccess"
	idataaccess "auth/dataaccess/interfaces"
	"auth/helpers"
	"auth/usecases"
	iusecases "auth/usecases/interfaces"

	"go.mongodb.org/mongo-driver/mongo"
)

type UserFactory struct {
	mongoCli *mongo.Client
	database string
	helpers  helpers.Helpers
}

func NewUserFactory(mongoCli *mongo.Client, database string, helpers helpers.Helpers) *UserFactory {
	return &UserFactory{mongoCli: mongoCli, database: database, helpers: helpers}
}

func (userFactory *UserFactory) GetUserController() *controllers.UserController {
	userRepository := InitializeUserRepository(userFactory.mongoCli, userFactory.database)
	userUseCase := InitializeUserUseCase(userRepository, userFactory.helpers)
	userController := InitializeUserController(userUseCase)
	return userController
}

func InitializeUserRepository(mongoCli *mongo.Client, database string) idataaccess.UserRepository {
	userRepo := dataaccess.NewUserMongoRepo(mongoCli, database)
	var userRepository idataaccess.UserRepository = userRepo
	return userRepository
}

func InitializeUserUseCase(UserRepository idataaccess.UserRepository, helpers helpers.Helpers) *usecases.UserUseCase {
	return usecases.NewUserUseCase(UserRepository, helpers)
}

func InitializeUserController(UserUseCase iusecases.UserUseCase) *controllers.UserController {
	return controllers.NewUserController(UserUseCase)
}
