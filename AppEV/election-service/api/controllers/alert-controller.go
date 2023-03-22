package controllers

import (
	iusecases "election-service/usecases/interfaces"

	"github.com/gofiber/fiber/v2"
)

type AlertController struct {
	alertUsecase iusecases.AlertUseCase
}

func NewAlertController(alertUsecase iusecases.AlertUseCase) *AlertController {
	return &AlertController{alertUsecase: alertUsecase}
}

func (controller *AlertController) GetAlertConfiguration(c *fiber.Ctx) error {
	electionId := c.Params("electionId")

	alertConfig, err := controller.alertUsecase.GetAlertConfiguration(electionId)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code":  fiber.StatusBadRequest,
			"Message":      err.Error(),
			"Alert Config": nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code":  fiber.StatusOK,
		"Message":      "Petición procesada correctamente",
		"Alert Config": alertConfig,
	})
}
