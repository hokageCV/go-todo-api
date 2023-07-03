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
	taskTitle := c.Body()

	err := storage.InsertIntoDB(string(taskTitle))
	if err != nil {
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
	title := c.Body()

	err := storage.UpdateTitleInDB(id, string(title))
	if err != nil {
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
