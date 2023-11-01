package validator

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

func BindAndCheck(c *fiber.Ctx, data any) error {
	if err := c.BodyParser(data); err != nil {
		return err
	}
	if v := validate.Struct(data); !v.Validate() {
		return v.Errors
	}

	return nil
}
