package main

import (
	"beexamxchat/domain"
	"beexamxchat/internal/api"
	"beexamxchat/internal/repository"
	"beexamxchat/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	allData := domain.AllData

	dataRepository := repository.NewDataRepository(allData)

	dataService := service.NewDataService(dataRepository)

	app := fiber.New()
	app.Use(logger.New())

	api.NewData(app, dataService)

	_ = app.Listen(":8080")
}
