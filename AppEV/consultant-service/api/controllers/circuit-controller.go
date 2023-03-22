package controllers

import (
	"consultant-service/models/read"
	iusecases "consultant-service/usecases/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
)

type CircuitController struct {
	circuitUseCase iusecases.CircuitUseCase
}

func NewCircuitController(circuitUseCase iusecases.CircuitUseCase) *CircuitController {
	return &CircuitController{circuitUseCase: circuitUseCase}
}

func (controller *CircuitController) GetVoteCoveragePerCircuit(c *fiber.Ctx) error {
	queryRequestTimeStamp := time.Now()
	var queryResponseTimeStamp time.Time

	var circuitCoverage []read.CircuitVoteCoverage

	electionId := c.Params("electionId")

	circuitCoverage, err := controller.circuitUseCase.GetVoteCoveragePerCircuit(electionId)

	if err != nil {
		queryResponseTimeStamp = time.Now()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"StatusCode":               fiber.StatusNotFound,
			"Message":                  err.Error(),
			"Circuit Coverage":         nil,
			"Query Request TimeStamp":  queryRequestTimeStamp,
			"Query Response TimeStamp": queryResponseTimeStamp,
			"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
		})
	}

	queryResponseTimeStamp = time.Now()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"StatusCode":               fiber.StatusOK,
		"Message":                  "Request processed correctly",
		"Circuit Coverage":         circuitCoverage,
		"Query Request TimeStamp":  queryRequestTimeStamp,
		"Query Response TimeStamp": queryResponseTimeStamp,
		"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
	})
}
