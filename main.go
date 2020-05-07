package main

import (
	"fmt"

	"github.com/JMercie/restApi_with_fiber/book"
	"github.com/JMercie/restApi_with_fiber/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := fiber.New()
	intitDatabase()
	defer database.DBConn.Close()

	setupRoutes(app)

	app.Listen(3000)
}

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBooks)
	app.Delete("/api/v1/book/:id", book.DeleteBooks)
}

func intitDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to the DB")
	}
	fmt.Println("DB succesfully connected")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("DB Migrated")
}
