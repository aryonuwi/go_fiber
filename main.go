package main

import (
	"github.com/aryonuwi/go_fiber.git/route"
	"github.com/gofiber/fiber/v2"
)

func main() {
	App := fiber.New()

	// Route Initial
	route.RouteInit(App)

	App.Listen(":3000")
}
