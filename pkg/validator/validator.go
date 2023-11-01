package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/kalougata/mall/pkg/e"
)

func BindAndCheck(c *fiber.Ctx, data any) error {
	if err := c.BodyParser(data); err != nil {
		return e.ErrInvalidRequestBody().WithErr(err)
	}
	if v := validate.Struct(data); !v.Validate() {
		return e.ErrBadRequest().WithErr(v.Errors)
	}

	return nil
}
