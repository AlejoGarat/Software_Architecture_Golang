package controllers

import (
	"consultant-service/models/write"
	iusecases "consultant-service/usecases/interfaces"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type AlertController struct {
	alertUsecase iusecases.AlertUseCase
}

func NewAlertController(alertUsecase iusecases.AlertUseCase) *AlertController {
	return &AlertController{alertUsecase: alertUsecase}
}

func (controller *AlertController) ModifyAlertConfiguration(c *fiber.Ctx) error {
	var alertConfig write.AlertConfiguration

	err := json.Unmarshal(c.Body(), &alertConfig)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	err = controller.alertUsecase.ModifyAlertConfiguration(alertConfig)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code": fiber.StatusOK,
		"Message":     "Valores esperados correctamente actualizados",
	})
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
