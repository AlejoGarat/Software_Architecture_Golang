package middlewares

import (
	"encoding/json"
	"votation-service/helpers"
	"votation-service/models/write"
	"votation-service/pipe"

	"github.com/gofiber/fiber/v2"
)

type VoteIssuance struct {
	pipeline  pipe.Pipeline
	logHelper helpers.LogHelper
}

func NewVoteIssuanceMiddleware(pipeline pipe.Pipeline, logHelper helpers.LogHelper) VoteIssuance {
	return VoteIssuance{pipeline: pipeline, logHelper: logHelper}
}

func (voteIssuanceMiddleware *VoteIssuance) VoteIssuanceMiddleware(c *fiber.Ctx) error {
	var vote write.Vote
	json.Unmarshal(c.Body(), &vote)
	err := voteIssuanceMiddleware.ExecuteVoteIssuanceFilters(vote)

	if err != nil {
		log := write.LoggingModel{Type: "Error", Operation: "Add Vote", Actor: "Voter", Description: err.Error()}
		voteIssuanceMiddleware.logHelper.SendLog(log)

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": 400,
			"Message":     err.Error(),
		})
	}

	return c.Next()
}

func (voteIssuanceMiddleware *VoteIssuance) ExecuteVoteIssuanceFilters(vote write.Vote) error {
	outChannelValue := voteIssuanceMiddleware.pipeline.ExecuteFilters(vote)
	return outChannelValue
}
