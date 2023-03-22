package routes

import (
	controllers "auth/api/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

type VoteIssuanceMiddleware func(c *fiber.Ctx) error

func PublicRoutes(a *fiber.App, userController *controllers.UserController) {

	route := a.Group("/identity-provider-api/v1")

	route.Post("/login", userController.Login)
}
