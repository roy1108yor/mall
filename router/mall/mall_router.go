package mallrouter

import (
	"github.com/gofiber/fiber/v2"
	mallv1 "github.com/kalougata/mall/api/v1/mall"
)

type MallHTTPServer *fiber.App

func NewMallHTTPServer(
	apiRouter *mallv1.MallAPIRouter,
) MallHTTPServer {
	app := fiber.New()

	v1Group := app.Group("/api/v1/mall")

	apiRouter.RegisterGuestAPIRouter(v1Group)

	return app
}
