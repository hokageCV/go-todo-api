package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/hokageCV/go-todo-api/api"
	"github.com/hokageCV/go-todo-api/storage"
	"github.com/hokageCV/go-todo-api/utils"
)

func main() {
	utils.InitializeLogger()

	err := storage.InitializeDB()
	if err != nil {
		panic(err)
	}

	db := storage.GetDB()
	defer db.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	api.SetupRoutes(app)

	app.Listen(":3000")
}
