package adminctrl

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/model"
	"github.com/kalougata/mall/pkg/e"
	"github.com/kalougata/mall/pkg/response"
	adminsrv "github.com/kalougata/mall/service/admin"
)

type umsResourceController struct {
	service adminsrv.UmsResourceService
}

type UmsResourceController interface {
	AddResourceCategory(c *fiber.Ctx) error
	AddResource(c *fiber.Ctx) error
}

// AddResource implements UmsResourceController.
func (rc *umsResourceController) AddResource(c *fiber.Ctx) error {
	data := &model.UmsResourceInReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}
	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}
	if err := rc.service.AddResource(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

// AddResourceCategory implements UmsResourceController.
func (rc *umsResourceController) AddResourceCategory(c *fiber.Ctx) error {
	data := &model.UmsResourceCategoryInReq{}
	if err := c.BodyParser(data); err != nil {
		return response.Build(c, e.ErrBadRequest().WithMsg(err.Error()), nil)
	}
	if v := validate.Struct(data); !v.Validate() {
		return response.Build(c, e.ErrInvalidRequestBody().WithErr(v.Errors), nil)
	}
	if err := rc.service.AddRecourceCategory(c.Context(), data); err != nil {
		return response.Build(c, err, nil)
	}

	return response.Build(c, nil, data)
}

func NewUmsResourceController(service adminsrv.UmsResourceService) UmsResourceController {
	return &umsResourceController{service}
}
