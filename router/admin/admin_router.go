package adminrouter

import (
	"github.com/gofiber/fiber/v2"
	adminv1 "github.com/kalougata/mall/api/v1/admin"
)

type AdminHTTPServer *fiber.App

func NewAdminHTTPServer(
	apiRouter *adminv1.AdminAPIRouter,
) AdminHTTPServer {
	app := fiber.New()

	v1Group := app.Group("/api/v1/admin")

	apiRouter.RegisterGuestAPIRouter(v1Group)

	return app
}