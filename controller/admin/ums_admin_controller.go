package adminctrl

import adminsrv "github.com/kalougata/mall/service/admin"

type umsAdminController struct {
	service adminsrv.UmsAdminService
}

type UmsAdminController interface {
}

func NewUmsAdminController(service adminsrv.UmsAdminService) UmsAdminController {
	return &umsAdminController{service}
}
