package adminctrl

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/model"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsAdminController struct {
	service adminsrv.UmsAdminService
}

type UmsAdminController interface {
	UmsAdminLogin(c *fiber.Ctx) error
	UmsAdminRegister(c *fiber.Ctx) error
}

// UmsAdminLogin implements UmsAdminController.
func (*umsAdminController) UmsAdminLogin(c *fiber.Ctx) error {
	data := &model.UmsAdminLoginReq{}
	if err := c.BodyParser(data); err != nil {
		return err
	}
	if v := validate.Struct(data); !v.Validate() {
		return v.Errors
	}

	data.LoginIpAddr = c.IP()
	data.LoginTime = time.Now()

	return c.SendString("ok!")
}

// UmsAdminRegister implements UmsAdminController.
func (*umsAdminController) UmsAdminRegister(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewUmsAdminController(service adminsrv.UmsAdminService) UmsAdminController {
	return &umsAdminController{service}
}
