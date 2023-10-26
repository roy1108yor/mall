package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world")
	})

	if err := app.Listen(":3000"); err != nil {
		log.Errorf("failed to start server: %s", err)
		return
	}

}
