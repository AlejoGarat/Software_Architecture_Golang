package routes

import (
	"election-service/api/common"
	"election-service/api/middlewares"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func PublicRoutes(a *fiber.App, controllers *common.Controllers) {
	route := a.Group("/electoral-api/v1")
	key := os.Getenv("KEY")

	a.Use(
		jwtware.New(jwtware.Config{SigningKey: []byte(key)}),
		middlewares.AuthorizationFilter,
	)

	route.Post("/election", controllers.ElectionController.AddElection)

	route.Get("/alert-configuration/:electionId", controllers.AlertController.GetAlertConfiguration)
	route.Get("/frequent-schedules/:electionId", controllers.ScheduleController.GetFrequentSchedules)
	route.Get("/department/coverage/:electionId", controllers.DepartmentController.GetDepartamenteCoverage)
	route.Get("/circuit/coverage/:electionId", controllers.CircuitController.GetVoteCoveragePerCircuit)
	route.Get("/votes/schedules", controllers.VoteController.GetVotersVoteSchedules)
	route.Get("/election/result/:electionId", controllers.ElectionController.GetElectionResult)
}
