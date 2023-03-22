package routes

import (
	controllers "uruguayanelectoralauthorityservice/api/controllers"

	fiber "github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App, controller *controllers.DataController) {
	route := a.Group("/api/autoridad-electoral-uruguay")

	route.Get("/data", controller.GetData)
}
