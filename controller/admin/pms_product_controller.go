package adminctrl

import "github.com/gofiber/fiber/v2"

type pmsProductController struct {
}

type PmsProductController interface {
	Create(c *fiber.Ctx) error
}

// Create implements PmsProductController.
func (*pmsProductController) Create(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewPmsProductController() PmsProductController {
	return &pmsProductController{}
}
