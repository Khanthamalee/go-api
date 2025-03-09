package crud

import "github.com/gofiber/fiber/v2"

func renderTemplate(c *fiber.Ctx) error {
	return c.Render("template", fiber.Map{
		"Name": "World",
	})
}
