package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterRoleAPIRouter(r fiber.Router) {
	r.Post("/add", api.umsRoleController.AddRole)
	r.Delete("/delete", api.umsRoleController.DeleteRole)
	r.Patch("/update", api.umsRoleController.UpdateRole)

	// 给角色分配菜单
	r.Post("/allocMenu", api.umsRoleController.AllocMenu)
}
