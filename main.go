package main

import (
	"go_fiber/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	App := fiber.New()

	// Route Initial
	routes.RouteInit(App)

	App.Listen(":3000")
}
