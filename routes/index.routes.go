package routes

import "github.com/gofiber/fiber/v2"

func RouteInit(r *fiber.App) {
	r.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"greatings": "Halo, Go Fiber Library"})
	})
}
