package controllers

import (
	"consultant-service/models/read"
	iusecases "consultant-service/usecases/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
)

type DepartmentController struct {
	departmentUseCase iusecases.DepartmentUseCase
}

func NewDepartmentController(departmentUseCase iusecases.DepartmentUseCase) *DepartmentController {
	return &DepartmentController{departmentUseCase: departmentUseCase}
}

func (controller *DepartmentController) GetDepartamenteCoverage(c *fiber.Ctx) error {
	queryRequestTimeStamp := time.Now()
	var queryResponseTimeStamp time.Time

	var departmentCoverage []read.DepartmentVoteCoverage

	electionId := c.Params("electionId")

	departmentCoverage, err := controller.departmentUseCase.GetVoteCoveragePerDepartment(electionId)

	if err != nil {
		queryResponseTimeStamp = time.Now()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"StatusCode":               fiber.StatusNotFound,
			"Message":                  err.Error(),
			"Department Coverage":      nil,
			"Query Request TimeStamp":  queryRequestTimeStamp,
			"Query Response TimeStamp": queryResponseTimeStamp,
			"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
		})
	}

	queryResponseTimeStamp = time.Now()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"StatusCode":               fiber.StatusOK,
		"Message":                  "Request processed correctly",
		"Department Coverage":      departmentCoverage,
		"Query Request TimeStamp":  queryRequestTimeStamp,
		"Query Response TimeStamp": queryResponseTimeStamp,
		"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
	})
}
