package controllers

import (
	"go-api-starter/internal/delivery/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type UserController struct {
	Log     *logrus.Logger
	UseCase *usecase.UserUseCase
}

func NewUserController(useCase *usecase.UserUseCase, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	return nil
}

func (c *UserController) Login(ctx *fiber.Ctx) error {
	return nil
}

func (c *UserController) Current(ctx *fiber.Ctx) error {
	return nil
}
