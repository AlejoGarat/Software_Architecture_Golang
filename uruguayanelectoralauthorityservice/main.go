package main

import (
	configs "uruguayanelectoralauthorityservice/api/configs"
	controllers "uruguayanelectoralauthorityservice/api/controllers"
	routes "uruguayanelectoralauthorityservice/api/routes"

	fiber "github.com/gofiber/fiber/v2"
)

func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	app := fiber.New(config)

	controller := controllers.NewDataController()

	routes.PublicRoutes(app, controller)

	app.Listen(":8081")
}
