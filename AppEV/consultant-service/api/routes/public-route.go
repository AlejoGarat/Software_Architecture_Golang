package routes

import (
	controllers "consultant-service/api/controllers"
	"consultant-service/api/middlewares"
	"os"

	fiber "github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

func PublicRoutes(a *fiber.App, filterController *controllers.FilterController,
	alertController *controllers.AlertController,
	circuitController *controllers.CircuitController,
	departmentController *controllers.DepartmentController,
	scheduleController *controllers.ScheduleController) {

	route := a.Group("/consultant-api/v1")
	key := os.Getenv("KEY")

	a.Use(
		jwtware.New(jwtware.Config{SigningKey: []byte(key)}),
		middlewares.AuthorizationFilter,
	)

	route.Put("/filters/election-beginning", filterController.ModifyElectionBeginningFilters)
	route.Put("/filters/election-end", filterController.ModifyElectionEndFilters)
	route.Put("/filters/vote-issuance", filterController.ModifyVoteIssuanceFilters)
	route.Put("/alert-configuration", alertController.ModifyAlertConfiguration)

	route.Get("/alert-configuration/:electionId", alertController.GetAlertConfiguration)
	route.Get("/frequent-schedules/:electionId", scheduleController.GetFrequentSchedules)
	route.Get("/department/coverage/:electionId", departmentController.GetDepartamenteCoverage)
	route.Get("/circuit/coverage/:electionId", circuitController.GetVoteCoveragePerCircuit)

}
