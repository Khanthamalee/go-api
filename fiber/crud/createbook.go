package crud

import "github.com/gofiber/fiber/v2"

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	book.ID = len(books) + 1
	books = append(books, *book)
	return c.JSON(books)
}
