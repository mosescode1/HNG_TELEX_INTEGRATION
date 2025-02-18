package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/data"
)



func HandleHome(c *fiber.Ctx) error {
	return c.Status(200).SendString("Welcome to Fiber Logs")
}

func HandleIntegrationJson(c *fiber.Ctx) error {
	data := data.GetIntegrationJson()
	return c.Status(200).JSON(data)
}