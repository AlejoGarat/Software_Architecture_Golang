package middlewares

import (
	"os"
	"votation-service/dataaccess"
	idataaccess "votation-service/dataaccess/interfaces"
	"votation-service/datasources"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userKey             = "user"
	idKey               = "id"
	consultingAgentRole = "voter"
)

func AuthorizationFilter(c *fiber.Ctx) error {
	user := c.Locals(userKey).(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims[idKey].(string)
	userToken := user.Raw

	mongoAddress := os.Getenv("MONGO_CLIENT")

	usersDB := os.Getenv("USERS_DB")

	mongoCli, err := getMongoClient(mongoAddress)

	if err != nil {
		return err
	}

	userRepo := dataaccess.NewUserMongoRepo(mongoCli, usersDB)
	var userRepository idataaccess.UserRepository = userRepo

	userRole, err := userRepository.GetUserRole(userId, userToken)

	if err != nil || userRole != consultingAgentRole {
		return c.Status(fiber.ErrUnauthorized.Code).JSON(fiber.Map{
			"StatusCode": fiber.ErrUnauthorized.Code,
			"Message":    "You have not authorization to access this resource",
		})
	}

	return c.Next()
}

func getMongoClient(mongoAddress string) (*mongo.Client, error) {
	return datasources.NewMongoDataSource(mongoAddress)
}
