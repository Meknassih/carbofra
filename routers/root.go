package routers

import (
	"github.com/gofiber/fiber/v2"
)

func Root(app *fiber.App) error {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{"Name": "Zaza"}) 
	})
	return nil
}
