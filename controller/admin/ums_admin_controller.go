package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsAdminController struct {
	service adminsrv.UmsAdminService
}

type UmsAdminController interface {
	UmsAdminLogin(c *fiber.Ctx) error
}

// UmsAdminLogin implements UmsAdminController.
func (*umsAdminController) UmsAdminLogin(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUmsAdminController(service adminsrv.UmsAdminService) UmsAdminController {
	return &umsAdminController{service}
}
