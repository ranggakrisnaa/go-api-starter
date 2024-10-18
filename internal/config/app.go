package config

import (
	"go-api-starter/internal/delivery/http/controllers"
	"go-api-starter/internal/delivery/http/route"
	"go-api-starter/internal/delivery/repository"
	"go-api-starter/internal/delivery/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	userRepository := repository.NewUserRepository(config.DB, config.Log)

	userUseCase := usecase.NewUserUseCase(config.DB, config.Log, config.Validate, userRepository)

	userController := controllers.NewUserController(userUseCase, config.Log)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}
	routeConfig.Setup()
}
