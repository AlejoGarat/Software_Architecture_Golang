package controllers

import (
	"consultant-service/models/read"
	iusecases "consultant-service/usecases/interfaces"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type FilterController struct {
	filterUsecase iusecases.FilterUseCase
}

func NewFilterController(filterUsecase iusecases.FilterUseCase) *FilterController {
	return &FilterController{filterUsecase: filterUsecase}
}

func (controller *FilterController) ModifyElectionBeginningFilters(c *fiber.Ctx) error {
	var filters read.ElectionBeginningFilters

	err := json.Unmarshal(c.Body(), &filters)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	err = controller.filterUsecase.ModifyElectionBeginningFilters(filters)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code": fiber.StatusOK,
		"Message":     "Filtros correctamente actualizados",
	})
}

func (controller *FilterController) ModifyElectionEndFilters(c *fiber.Ctx) error {
	var filters read.ElectionEndFilters

	err := json.Unmarshal(c.Body(), &filters)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	err = controller.filterUsecase.ModifyElectionEndFilters(filters)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code": fiber.StatusOK,
		"Message":     "Filtros correctamente actualizados",
	})
}

func (controller *FilterController) ModifyVoteIssuanceFilters(c *fiber.Ctx) error {
	var filters read.VoteIssuanceFilters

	err := json.Unmarshal(c.Body(), &filters)
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	err = controller.filterUsecase.ModifyVoteIssuanceFilters(filters)

	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"Status Code": fiber.StatusOK,
		"Message":     "Filtros correctamente actualizados",
	})
}
