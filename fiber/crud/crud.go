package crud

import "github.com/gofiber/fiber/v2"

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book = []Book{
	{ID: 1, Title: "1984", Author: "George Orwell"},
	{ID: 2, Title: "The Great Gatsby", Author: "F.SCott Fitzgerald"},
}

func Callcrud() {
	app := fiber.New()
	app.Get("/books", getBooks)
	app.Get("/books/:id", getBook)
	app.Post("/books/", createBook)
	app.Post("/books/:id", updateBook)
	app.Post("/books/:id", deleteBook)
	app.Listen(":8080")

}
