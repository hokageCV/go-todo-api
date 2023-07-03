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

func InitializeDB() {
	logger := utils.GetLogger()
	connStr := "user=CBlizzard password=AGxnpgCy65PK dbname=neondb host=ep-sparkling-snow-514514-pooler.us-east-2.aws.neon.tech sslmode=verify-full"

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("can't connect to DB😱 ", zap.Error(err))
	}

	err = createTableIfNotExists()
	if err != nil {
		logger.Error("can't create table😱 ", zap.Error(err))
	}

}

func GetDB() *sql.DB {
	return db
}

func createTableIfNotExists() error {

	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		is_done BOOLEAN
	)
`)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}
