package middlewares

import (
	"analytics/dataaccess"
	idataaccess "analytics/dataaccess/interfaces"
	"analytics/datasources"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	consultingAgentRole = "consulting_agent"
)

func AuthorizationFilter(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	mongoAddress := os.Getenv("MONGO_CLIENT")

	usersDB := os.Getenv("USERS_DB")

	mongoCli, err := getMongoClient(mongoAddress)

	if err != nil {
		return err
	}

	userRepo := dataaccess.NewUserMongoRepo(mongoCli, usersDB)
	var userRepository idataaccess.UserRepository = userRepo

	userRole, err := userRepository.GetUserRole(userId)

	if err != nil {
		return err
	}

	if userRole != consultingAgentRole {
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
