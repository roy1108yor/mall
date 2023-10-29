package adminv1

import adminctrl "github.com/kalougata/mall/controller/admin"

type AdminAPIRouter struct {
	umsAdminController adminctrl.UmsAdminController
	umsRoleController  adminctrl.UmsRoleController
	umsMenuController  adminctrl.UmsMenuController
}

func NewAdminAPIRouter(
	umsAdminController adminctrl.UmsAdminController,
	umsRoleController adminctrl.UmsRoleController,
	umsMenuController adminctrl.UmsMenuController,
) *AdminAPIRouter {
	return &AdminAPIRouter{
		umsAdminController,
		umsRoleController,
		umsMenuController,
	}
}
