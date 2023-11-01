package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/response"
	"github.com/kalougata/mall/pkg/validator"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsMenuController struct {
	service adminsrv.UmsMenuService
}

type UmsMenuController interface {
	AddMenu(c *fiber.Ctx) error
	TreeList(c *fiber.Ctx) error
}

// TreeList implements UmsMenuController.
func (mc *umsMenuController) TreeList(c *fiber.Ctx) error {
	treeList, err := mc.service.TreeList(c.Context())
	if err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, treeList)
}

// AddMenu implements UmsMenuController.
func (mc *umsMenuController) AddMenu(c *fiber.Ctx) error {
	data := &model.UmsMenuInReq{}
	if err := validator.BindAndCheck(c, data); err != nil {
		return response.Build(c, err, nil)
	}
	if err := mc.service.AddMenu(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

func NewUmsMenuController(service adminsrv.UmsMenuService) UmsMenuController {
	return &umsMenuController{service}
}
