package app

import (
	"go-template-clean/config"
	"go-template-clean/pkg/cache"
	"go-template-clean/pkg/db"
	"go-template-clean/pkg/logger"

	"github.com/gofiber/fiber/v2"

	exampleusecasehandler "go-template-clean/internal/usecase/exampleusecase/controller/http"
	exampleusecaserepositorydb "go-template-clean/internal/usecase/exampleusecase/repository"
	exampleusecaseservice "go-template-clean/internal/usecase/exampleusecase/service"
)

func NewApplication(app fiber.Router, logger logger.Logger, db *db.DB, cache cache.Engine, config *config.Configuration) {

	exampleRepository := exampleusecaserepositorydb.NewExampleRepository(db, logger, config)
	exampleService := exampleusecaseservice.NewExampleService(exampleRepository, config, cache, logger)
	exampleHandler := exampleusecasehandler.NewExampleHandler(exampleService, config)
	exampleHandler.InitRoute(app)
}
