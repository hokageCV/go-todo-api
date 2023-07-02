package storage

import (
	"database/sql"

	"github.com/hokageCV/go-todo-api/utils"
)

func InitializeDB() (*sql.DB, error) {
	connStr := "user=CBlizzard password=AGxnpgCy65PK dbname=neondb host=ep-sparkling-snow-514514-pooler.us-east-2.aws.neon.tech sslmode=verify-full"

	db, err := sql.Open("postgres", connStr)
	utils.CheckNilErr(err)

	err = db.Ping()
	utils.CheckNilErr(err)

	return db, nil
}
