package crud

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func getBook(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params(("id")))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	for _, book := range books {
		if book.ID == id {
			return c.JSON(book)
		}
	}
	return c.SendStatus(fiber.StatusBadRequest)
}
