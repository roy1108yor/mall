package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsRoleController struct {
	service adminsrv.UmsRoleService
}

type UmsRoleController interface {
	AddRole(c *fiber.Ctx) error
}

// AddRole 添加角色
func (rc *umsRoleController) AddRole(c *fiber.Ctx) error {
	data := &model.UmsRoleReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}
	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}

	rc.service.AddRole(c.Context(), data)

	return nil
}

func NewUmsRoleController(service adminsrv.UmsRoleService) UmsRoleController {
	return &umsRoleController{service}
}
