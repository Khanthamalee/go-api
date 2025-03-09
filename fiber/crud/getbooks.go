package crud

import "github.com/gofiber/fiber/v2"

func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}
