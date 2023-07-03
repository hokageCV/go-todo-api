package storage

import (
	"github.com/hokageCV/go-todo-api/types"
	"github.com/hokageCV/go-todo-api/utils"
	"go.uber.org/zap"
)

func InsertIntoDB(title string) {
	db := GetDB()
	logger := utils.GetLogger()

	stmt, err := db.Prepare("INSERT INTO todos (title, is_done) VALUES ($1, $2)")
	if err != nil {
		logger.Error("🛑 can't prepare statement ", zap.Error(err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, false)
	if err != nil {
		logger.Error("🛑 can't execute statement ", zap.Error(err))
	} else {
		logger.Info("✅ successfully inserted into table")
	}
}

func GetAllTodoFromDB() []types.Todo {

	db := GetDB()
	logger := utils.GetLogger()

	rows, err := db.Query("SELECT * FROM todos")
	if err != nil {
		logger.Error("🛑 can't query table ", zap.Error(err))
	}
	defer rows.Close()

	var todoList []types.Todo

	for rows.Next() {
		var task types.Todo
		err = rows.Scan(&task.ID, &task.Title, &task.IsDone)
		if err != nil {
			logger.Error("🛑 can't scan row ", zap.Error(err))
		}
		todoList = append(todoList, task)
	}
	logger.Info("✅ successfully fetched all rows from table")

	return todoList
}

func UpdateInDB(taskID string) {
	db := GetDB()
	logger := utils.GetLogger()

	stmt, err := db.Prepare("UPDATE todos SET is_done = true WHERE id = $1")
	if err != nil {
		logger.Error("🛑 can't prepare statement ", zap.Error(err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskID)
	if err != nil {
		logger.Error("🛑 can't execute statement ", zap.Error(err))
	} else {
		logger.Info("✅ successfully updated row in table")
	}
}

func DeleteFromDB(taskID string) {
	db := GetDB()
	logger := utils.GetLogger()

	stmt, err := db.Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		logger.Error("🛑 can't prepare statement ", zap.Error(err))
	}
	defer stmt.Close()

	_, err = stmt.Exec(taskID)
	if err != nil {
		logger.Error("🛑 can't execute statement ", zap.Error(err))
	} else {
		logger.Info("✅ successfully deleted row in table")
	}
}
