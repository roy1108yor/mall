package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterGuestAPIRouter(r fiber.Router) {
	r.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"msg": "pong!",
		})
	})
}
