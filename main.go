package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hokageCV/go-todo-api/storage"
)

func main() {
	storage.InitializeDB()
	db := storage.GetDB()
	defer db.Close()

	app := fiber.New()

	// routing

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
