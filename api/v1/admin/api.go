package adminv1

import adminctrl "github.com/kalougata/mall/controller/admin"

type AdminAPIRouter struct {
	umsAdminController adminctrl.UmsAdminController
	umsRoleController  adminctrl.UmsRoleController
}

func NewAdminAPIRouter(
	umsAdminController adminctrl.UmsAdminController,
	umsRoleController adminctrl.UmsRoleController,
) *AdminAPIRouter {
	return &AdminAPIRouter{
		umsAdminController,
		umsRoleController,
	}
}
