package adminv1

import adminctrl "github.com/kalougata/mall/controller/admin"

type AdminAPIRouter struct {
	umsAdminController adminctrl.UmsAdminController
}

func NewAdminAPIRouter(
	umsAdminController adminctrl.UmsAdminController,
) *AdminAPIRouter {
	return &AdminAPIRouter{
		umsAdminController,
	}
}
