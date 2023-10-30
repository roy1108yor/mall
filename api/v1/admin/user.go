package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterAdminUserAPIRouter(r fiber.Router) {
	r.Post("allocRole", api.umsAdminController.AllocRoleForAdmin)
}
