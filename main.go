package main

import (
	"github.com/aryonuwi/go_fiber.git/conf"
	"github.com/aryonuwi/go_fiber.git/route"
	"github.com/gofiber/fiber/v2"
)

func main() {

	conf.DatabaseInit()
	App := fiber.New()

	// Route Initial
	route.RouteInit(App)
	route.RouteUsers(App)

	App.Listen(":3000")
}
