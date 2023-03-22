package controllers

import (
	"election-service/models/read"
	iusecases "election-service/usecases/interfaces"
	"encoding/json"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type VoteController struct {
	voteUseCase iusecases.VoteUseCase
}

func NewVoteController(voteUseCase iusecases.VoteUseCase) *VoteController {
	return &VoteController{voteUseCase: voteUseCase}
}

func (controller *VoteController) GetVotersVoteSchedules(c *fiber.Ctx) error {
	queryRequestTimeStamp := time.Now()
	var queryResponseTimeStamp time.Time

	var request read.VoterSchedulesRequest

	err := json.Unmarshal(c.Body(), &request)

	if err != nil {
		queryResponseTimeStamp = time.Now()
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code":              fiber.ErrBadRequest.Code,
			"Message":                  err.Error(),
			"Query Request TimeStamp":  queryRequestTimeStamp,
			"Query Response TimeStamp": queryResponseTimeStamp,
			"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
		})
	}

	voterSchedules, err := controller.voteUseCase.GetVoterVotingSchedules(request.Election, request.VoterId)

	if err != nil {
		queryResponseTimeStamp = time.Now()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status Code":              fiber.ErrBadRequest.Code,
			"Message":                  "There are no votes registered with the document provided",
			"Schedules":                nil,
			"Query Request TimeStamp":  queryRequestTimeStamp,
			"Query Response TimeStamp": queryResponseTimeStamp,
			"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
		})
	}

	queryResponseTimeStamp = time.Now()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status Code":              fiber.StatusOK,
		"Message":                  "Request successfully completed",
		"Schedules":                voterSchedules,
		"Query Request TimeStamp":  queryRequestTimeStamp,
		"Query Response TimeStamp": queryResponseTimeStamp,
		"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).Seconds(),
	})
}
