package main

import (
	"github.com/hokageCV/go-todo-api/storage"
	"github.com/hokageCV/go-todo-api/types"
	"github.com/hokageCV/go-todo-api/utils"
	"go.uber.org/zap"
)

func main() {
	utils.InitializeLogger()
	storage.InitializeDB()
	Logger := utils.GetLogger()

	db := storage.GetDB()
	defer db.Close()

	task := types.Todo{
		Title:  "rasengan",
		IsDone: false,
	}

	// Execute the INSERT query
	_, err := db.Exec("INSERT INTO todos (title, is_done) VALUES ($1, $2)", task.Title, task.IsDone)
	if err != nil {
		Logger.Error("can't insert into table ", zap.Error(err))
	}

}
