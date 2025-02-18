package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mosescode1/data"
)



func HandleIntegrationJson(c *fiber.Ctx) error {
	data := data.GetIntegrationJson()
	return c.Status(200).JSON(data)
}