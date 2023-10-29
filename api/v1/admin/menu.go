package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterMenuAPIRouter(r fiber.Router) {
	r.Post("add")
}
