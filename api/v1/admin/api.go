package adminv1

import adminctrl "github.com/kalougata/mall/controller/admin"

type AdminAPIRouter struct {
	umsAdminController   adminctrl.UmsAdminController
	umsRoleController    adminctrl.UmsRoleController
	umsMenuController    adminctrl.UmsMenuController
	umsResourceControler adminctrl.UmsResourceController
	pmsProductController adminctrl.PmsProductController
}

func NewAdminAPIRouter(
	umsAdminController adminctrl.UmsAdminController,
	umsRoleController adminctrl.UmsRoleController,
	umsMenuController adminctrl.UmsMenuController,
	umsResourceControler adminctrl.UmsResourceController,
	pmsProductController adminctrl.PmsProductController,
) *AdminAPIRouter {
	return &AdminAPIRouter{
		umsAdminController,
		umsRoleController,
		umsMenuController,
		umsResourceControler,
		pmsProductController,
	}
}
