package controllers

import (
	"analytics/models/read"
	iusecases "analytics/usecases/interfaces"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ScheduleController struct {
	scheduleUseCase iusecases.ScheduleUseCase
}

func NewScheduleController(scheduleUseCase iusecases.ScheduleUseCase) *ScheduleController {
	return &ScheduleController{scheduleUseCase: scheduleUseCase}
}

func (controller *ScheduleController) GetFrequentSchedules(c *fiber.Ctx) error {
	queryRequestTimeStamp := time.Now()
	var queryResponseTimeStamp time.Time

	var schedules read.FrequentVotationSchedules

	electionId := c.Params("electionId")

	schedules, err := controller.scheduleUseCase.GetFrequentSchedules(electionId)

	if err != nil {
		queryResponseTimeStamp = time.Now()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"StatusCode":               fiber.StatusNotFound,
			"Message":                  err.Error(),
			"Frequent Schedules":       schedules,
			"Query Request TimeStamp":  queryRequestTimeStamp,
			"Query Response TimeStamp": queryResponseTimeStamp,
			"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp),
		})
	}

	queryResponseTimeStamp = time.Now()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"StatusCode":               fiber.StatusOK,
		"Message":                  "Request processed correctly",
		"Frequent Schedules":       schedules,
		"Query Request TimeStamp":  queryRequestTimeStamp,
		"Query Response TimeStamp": queryResponseTimeStamp,
		"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp),
	})
}
