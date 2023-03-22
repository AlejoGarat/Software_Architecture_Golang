package controllers

import (
	"election-service/models/read"
	iusecases "election-service/usecases/interfaces"
	"encoding/json"
	"time"

	fiber "github.com/gofiber/fiber/v2"
)

type ElectionController struct {
	electionUsecase iusecases.ElectionUsecase
}

func NewElectionController(elctionUsecase iusecases.ElectionUsecase) *ElectionController {
	return &ElectionController{electionUsecase: elctionUsecase}
}

func (controller *ElectionController) AddElection(c *fiber.Ctx) error {
	var election read.ElectionData

	err := json.Unmarshal(c.Body(), &election)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.ErrBadRequest,
			"Message":     err.Error(),
		})
	}

	err = controller.electionUsecase.AddElection(election)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.ErrBadRequest,
			"Message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code": fiber.StatusCreated,
		"Message":     "Election created succesfully",
	})
}

func (controller *ElectionController) GetElectionResult(c *fiber.Ctx) error {
	queryRequestTimeStamp := time.Now()
	var queryResponseTimeStamp time.Time

	var electionResult read.ElectionResult

	electionId := c.Params("electionId")

	electionResult, err := controller.electionUsecase.GetElectionResult(electionId)

	if err != nil {
		queryResponseTimeStamp = time.Now()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"Status Code":              fiber.StatusNotFound,
			"Message":                  err.Error(),
			"Election Result":          nil,
			"Query Request TimeStamp":  queryRequestTimeStamp,
			"Query Response TimeStamp": queryResponseTimeStamp,
			"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).String(),
		})
	}

	queryResponseTimeStamp = time.Now()

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code":              fiber.StatusCreated,
		"Message":                  "Request processed correctly",
		"Election Result":          electionResult,
		"Query Request TimeStamp":  queryRequestTimeStamp,
		"Query Response TimeStamp": queryResponseTimeStamp,
		"Query Processing Time":    queryResponseTimeStamp.Sub(queryRequestTimeStamp).String(),
	})
}
