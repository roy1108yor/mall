package adminv1

import "github.com/gofiber/fiber/v2"

func (api *AdminAPIRouter) RegisterResourceAPIRouter(r fiber.Router) {
	r.Post("add", api.umsResourceControler.AddResource)
	r.Post("category/add", api.umsResourceControler.AddResourceCategory)
}
