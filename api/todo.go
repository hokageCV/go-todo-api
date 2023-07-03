package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hokageCV/go-todo-api/storage"
	"github.com/hokageCV/go-todo-api/utils"
	"go.uber.org/zap"
)

func GetAllTodo(c *fiber.Ctx) error {
	logger := utils.GetLogger()

	todoList, err := storage.GetAllTodoFromDB()
	if err != nil {
		logger.Error("🛑 can't get all todo from db ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve todos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(todoList)
}

func CreateTodo(c *fiber.Ctx) error {
	logger := utils.GetLogger()

	var request struct {
		Title string `json:"title"`
	}

	err := c.BodyParser(&request)
	if err != nil {
		logger.Error("🛑 can't parse request body ", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	errr := storage.InsertIntoDB(request.Title)
	if errr != nil {
		logger.Error("🛑 can't insert into db ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create todo",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Todo created successfully",
	})
}

func MarkTodoDone(c *fiber.Ctx) error {
	logger := utils.GetLogger()
	id := c.Params("id")

	err := storage.UpdateDoneInDB(id)
	if err != nil {
		logger.Error("🛑 can't update todo in db ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update todo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo Done successfully",
	})
}

func UpdateTodoTitle(c *fiber.Ctx) error {
	logger := utils.GetLogger()
	id := c.Params("id")

	var request struct {
		Title string `json:"title"`
	}

	err := c.BodyParser(&request)
	if err != nil {
		logger.Error("🛑 can't parse request body ", zap.Error(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	errr := storage.UpdateTitleInDB(id, request.Title)
	if errr != nil {
		logger.Error("🛑 can't update todo in db ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update todo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo updated successfully",
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	logger := utils.GetLogger()
	id := c.Params("id")

	err := storage.DeleteFromDB(id)
	if err != nil {
		logger.Error("🛑 can't delete todo from db ", zap.Error(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete todo",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Todo deleted successfully",
	})
}
