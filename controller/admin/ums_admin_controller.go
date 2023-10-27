package adminctrl

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
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
func (uc *umsAdminController) UmsAdminLogin(c *fiber.Ctx) error {
	data := &model.UmsAdminLoginReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}
	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}

	data.LoginIpAddr = c.IP()
	data.LoginTime = time.Now()

	return response.Build(c, nil, "ok!")
}

// UmsAdminRegister implements UmsAdminController.
func (uc *umsAdminController) UmsAdminRegister(c *fiber.Ctx) error {
	data := &model.UmsAdminRegisterReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}

	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}

	data.RegIpAddr = c.IP()

	return response.Build(c, nil, "ok!")
}

func NewUmsAdminController(service adminsrv.UmsAdminService) UmsAdminController {
	return &umsAdminController{service}
}
