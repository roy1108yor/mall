package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterMenuAPIRouter(r fiber.Router) {
	r.Post("add", api.umsMenuController.AddMenu)
	r.Get("treeList", api.umsMenuController.TreeList)
}
