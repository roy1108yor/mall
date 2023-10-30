package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterProductAPIRouter(r fiber.Router) {
	r.Post("/create", api.pmsProductController.Create)
}
