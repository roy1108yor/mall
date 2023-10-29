package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
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
	treeList, err := mc.service.ListWithTree(c.Context())
	if err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, treeList)
}

// AddMenu implements UmsMenuController.
func (mc *umsMenuController) AddMenu(c *fiber.Ctx) error {
	data := &model.UmsMenuInReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}
	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}
	if err := mc.service.AddMenu(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

func NewUmsMenuController(service adminsrv.UmsMenuService) UmsMenuController {
	return &umsMenuController{service}
}
