package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterGuestAPIRouter(r fiber.Router) {
	r.Post("/login", api.umsAdminController.UmsAdminLogin)
	r.Post("/register", api.umsAdminController.UmsAdminRegister)
}
