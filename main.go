package main

import (
	"go_fiber/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	App := fiber.New()

	// Route Initial
	route.RouteInit(App)

	App.Listen(":3000")
}
