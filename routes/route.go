package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/handler"
)


func SetupRoutes(app *fiber.App) {
	app.Get("/integration_json", handler.HandleIntegrationJson)
}