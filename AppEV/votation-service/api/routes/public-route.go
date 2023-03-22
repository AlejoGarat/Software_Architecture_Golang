package routes

import (
	"os"
	controllers "votation-service/api/controllers"
	"votation-service/api/middlewares"

	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

type VoteIssuanceMiddleware func(c *fiber.Ctx) error

func PublicRoutes(a *fiber.App, voteController *controllers.VoteController, voteIssuance middlewares.VoteIssuance) {

	route := a.Group("/votation-api/v1")

	key := os.Getenv("KEY")

	a.Use(
		jwtware.New(jwtware.Config{SigningKey: []byte(key)}),
		middlewares.AuthorizationFilter,
	)

	route.Post("/mail", voteController.SendMail)

	a.Use(
		voteIssuance.VoteIssuanceMiddleware,
	)

	route.Post("/vote", voteController.AddVote)
}
