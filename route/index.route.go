package route

import "github.com/gofiber/fiber/v2"

func RouteInit(route *fiber.App) {
	route.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{"greatings": "Halo, Go Fiber Library"})
	})
}
