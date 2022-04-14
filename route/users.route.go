package route

import (
	"github.com/aryonuwi/go_fiber.git/handler"
	"github.com/gofiber/fiber/v2"
)

func RouteUsers(route *fiber.App) {

	route.Get("/users", handler.UsersHandler)

}
