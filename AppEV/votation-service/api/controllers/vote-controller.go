package controllers

import (
	"encoding/json"
	"strings"
	"time"
	"votation-service/models/read"
	"votation-service/models/write"
	iusecases "votation-service/usecases/interfaces"

	fiber "github.com/gofiber/fiber/v2"
)

const (
	datePos = 0
	hourPos = 1
)

type VoteController struct {
	voteUsecase iusecases.VoteUseCase
}

func NewVoteController(voteUsecase iusecases.VoteUseCase) *VoteController {
	return &VoteController{voteUsecase: voteUsecase}
}

func (controller *VoteController) AddVote(c *fiber.Ctx) error {
	var vote write.Vote

	date, hour := getDateAndHour()

	voteInfo := write.Info{Date: date, Hour: hour}

	err := json.Unmarshal(c.Body(), &vote)
	if err != nil {

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": 400,
			"Message":     err.Error(),
		})
	}

	vote.Info = voteInfo
	vote.StartingTime = time.Now()

	controller.voteUsecase.AddVote(vote)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code": 201,
		"Message":     "Your vote is being processed...",
	})
}

func (controller *VoteController) SendMail(c *fiber.Ctx) error {
	queryRequestTimeStamp := time.Now()

	var constancyRequest read.ConstancyRequest

	var constancyDBData write.ConstancyDBData

	err := json.Unmarshal(c.Body(), &constancyRequest)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": 400,
			"Message":     err.Error(),
		})
	}

	err = controller.voteUsecase.SendMailConstancy(constancyRequest)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": 400,
			"Message":     err.Error(),
		})
	}

	constancyDBData.VoterDocument = constancyRequest.VoterId
	constancyDBData.Timestamp = queryRequestTimeStamp

	controller.voteUsecase.UpdateConstancyDBData(constancyDBData, constancyRequest.ElectionId)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status Code": 200,
		"Message":     "Constancy is being prepared. You will be soon receiving a mail",
	})
}

func getDateAndHour() (string, string) {
	currentTime := time.Now()

	timeData := currentTime.Format("2006.01.02 15:00")

	splittedTimeData := strings.Split(timeData, " ")

	date := splittedTimeData[datePos]

	hour := splittedTimeData[hourPos]

	return date, hour
}
