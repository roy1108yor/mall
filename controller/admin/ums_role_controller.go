package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
	"github.com/kalougata/mall/pkg/validator"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsRoleController struct {
	service adminsrv.UmsRoleService
}

type UmsRoleController interface {
	AddRole(c *fiber.Ctx) error
	DeleteRole(c *fiber.Ctx) error
	UpdateRole(c *fiber.Ctx) error
	AllocMenuForRole(c *fiber.Ctx) error
}

// AllocMenu 为角色分配菜单
func (rc *umsRoleController) AllocMenuForRole(c *fiber.Ctx) error {
	data := &model.AllocMenuForRoleReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}

	if err := rc.service.AllocMenuForRole(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

// AddRole 添加角色
func (rc *umsRoleController) AddRole(c *fiber.Ctx) error {
	data := &model.AddUmsRoleReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}
	if err := rc.service.AddRole(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

// DeleteRole 批量删除角色
func (rc *umsRoleController) DeleteRole(c *fiber.Ctx) error {
	data := &model.DelUmsRoleReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}
	if len(data.Ids) <= 0 {
		return response.Build(c, e.ErrInvalidRequestBody().WithMsg("角色ID不能为空"), nil)
	}
	if err := rc.service.BatchDeleteRole(c.Context(), data.Ids); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

// UpdateROle 更新角色信息
func (rc *umsRoleController) UpdateRole(c *fiber.Ctx) error {
	data := &model.UpdateUmsRoleReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}
	if err := rc.service.UpdateRole(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

func NewUmsRoleController(service adminsrv.UmsRoleService) UmsRoleController {
	return &umsRoleController{service}
}
