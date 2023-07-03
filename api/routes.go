package api

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	apiRouter := app.Group("/api")

	apiRouter.Get("/todo", GetAllTodo)
	apiRouter.Post("/todo", CreateTodo)
	apiRouter.Patch("/todo/:id/done", MarkTodoDone)
	apiRouter.Patch("/todo/:id", UpdateTodoTitle)
	apiRouter.Delete("/todo/:id", DeleteTodo)
}
