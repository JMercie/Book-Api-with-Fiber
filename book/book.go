package book

import (
	"github.com/JMercie/restApi_with_fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DBConn

	var books []Book

	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var books Book

	db.Find(&books, id)
	c.JSON(books)
}

func NewBooks(c *fiber.Ctx) {
	db := database.DBConn

	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(&book)
	c.JSON(book)
}

func DeleteBooks(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("no book found with given id")
		return
	}

	db.Delete(&book)
	c.Send("book succesfully deleted from DB")
}
