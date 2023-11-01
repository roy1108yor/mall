package adminctrl

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
	"github.com/kalougata/mall/pkg/validator"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsAdminController struct {
	service adminsrv.UmsAdminService
}

type UmsAdminController interface {
	UmsAdminLogin(c *fiber.Ctx) error
	UmsAdminRegister(c *fiber.Ctx) error
	AllocRoleForAdmin(c *fiber.Ctx) error
}

// AllocRoleForAdmin 为用户分配角色
func (ac *umsAdminController) AllocRoleForAdmin(c *fiber.Ctx) error {
	data := &model.UmsRoleRelationInReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}
	if len(data.RoleIds) <= 0 {
		return response.Build(c, e.New(http.StatusBadRequest, "roleIds 必填项"), nil)
	}
	if err := ac.service.AllocRoleForAdmin(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, nil)
}

// UmsAdminLogin 管理员登录
func (uc *umsAdminController) UmsAdminLogin(c *fiber.Ctx) error {
	data := &model.UmsAdminLoginReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}

	data.LoginIpAddr = c.IP()
	data.LoginTime = time.Now()
	if resp, err := uc.service.Login(c.Context(), data); err == nil {
		return response.Build(c, nil, resp)
	} else {
		return response.Build(c, err, resp)
	}
}

// UmsAdminRegister 管理员注册
func (uc *umsAdminController) UmsAdminRegister(c *fiber.Ctx) error {
	data := &model.UmsAdminInReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(err), nil)
	}
	data.RegIpAddr = c.IP()
	if err := uc.service.Register(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, nil)
}

func NewUmsAdminController(service adminsrv.UmsAdminService) UmsAdminController {
	return &umsAdminController{service}
}
