package routes

import (
	"analytics/api/common"
	"analytics/api/middlewares"

	jwtware "github.com/gofiber/jwt/v3"

	fiber "github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App, controllers *common.Controllers) {
	route := a.Group("/analytics-api/v1")

	// JWT Middleware
	a.Use(
		jwtware.New(jwtware.Config{SigningKey: []byte("secret")}),
		middlewares.AuthorizationFilter,
	)

	route.Get("/frequent-schedules/:electionId", controllers.ScheduleController.GetFrequentSchedules)

	route.Get("/department/coverage/:electionId", controllers.DepartmentController.GetDepartamenteCoverage)

	route.Get("/circuit/coverage/:electionId", controllers.CircuitController.GetVoteCoveragePerCircuit)
}
