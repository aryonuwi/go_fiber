package handler

import "github.com/gofiber/fiber/v2"

func UsersHandler(ctx *fiber.Ctx) error {
	return ctx.JSON(fiber.Map{
		"greatings": "this from handeler",
	})
}
