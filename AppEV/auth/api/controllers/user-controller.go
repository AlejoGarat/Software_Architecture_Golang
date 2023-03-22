package controllers

import (
	"auth/models/read"
	iusecases "auth/usecases/interfaces"

	"github.com/gofiber/fiber/v2"

	"encoding/json"
)

type UserController struct {
	userUseCase iusecases.UserUseCase
}

func NewUserController(userUseCase iusecases.UserUseCase) *UserController {
	return &UserController{userUseCase: userUseCase}
}

func (controller *UserController) Login(c *fiber.Ctx) error {
	var user read.User

	requestErr := json.Unmarshal(c.Body(), &user)

	if requestErr != nil {

		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Status Code": fiber.ErrBadRequest.Code,
			"Message":     requestErr.Error(),
		})
	}

	hashedPsw := controller.userUseCase.HashPassword(user.Password)

	token, err := controller.userUseCase.Login(user.Id, hashedPsw)

	if err != nil {

		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Status Code": fiber.StatusBadRequest,
			"Message":     err.Error(),
		})
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"Status Code": fiber.StatusInternalServerError,
			"Message":     err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Status Code": fiber.StatusOK,
		"Message":     "Correctly logged in",
		"Token":       token,
	})
}
