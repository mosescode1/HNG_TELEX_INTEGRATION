package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/mosescode1/routes"
)


func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())

	routes.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}