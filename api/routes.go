package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// ================================
	todoRouter := app.Group("/api")

	todoRouter.Get("/todo", GetAllTodo)
	todoRouter.Post("/todo", CreateTodo)
	todoRouter.Patch("/todo/:id/done", MarkTodoDone)
	todoRouter.Patch("/todo/:id", UpdateTodoTitle)
	todoRouter.Delete("/todo/:id", DeleteTodo)

	// ================================
	miscRouter := app.Group("/misc")

	miscRouter.Get("/pokemon", GetRandomPokemon)
}
