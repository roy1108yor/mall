package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func Run(app *fiber.App, addr string) {
	if err := app.Listen(addr); err != nil {
		log.Panic(err)
	}
}
