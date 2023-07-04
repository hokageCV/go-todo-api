package storage

import (
	"database/sql"
	"fmt"

	"github.com/hokageCV/go-todo-api/utils"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

var db *sql.DB
var err error

func InitializeDB() error {
	logger := utils.GetLogger()
	connStr := utils.GetEnvVariable("CONENCTION_STRING")

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("can't connect to DB😱 ", zap.Error(err))
	}

	err = db.Ping()
	if err != nil {
		logger.Error("can't ping DB😱 ", zap.Error(err))
		return err
	}

	err = createTableIfNotExists()
	if err != nil {
		logger.Error("can't create table😱 ", zap.Error(err))
	}

	// // delete table
	// _, err = db.Exec("DROP TABLE todos")

	return nil

}

func GetDB() *sql.DB {
	return db
}

func createTableIfNotExists() error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id VARCHAR(255) PRIMARY KEY,
		title TEXT,
		is_done BOOLEAN
	)
`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
