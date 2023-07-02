package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CheckNilErr(err error, customMessage ...string) {
	if err != nil {
		if len(customMessage) > 0 {
			log.Fatal(customMessage[0])
		} else {
			log.Fatal(err)
		}
	}
}

func GetEnvVariable(key string) string {
	err := godotenv.Load(".env")
	CheckNilErr(err, "Error loading .env file")

	return os.Getenv(key)
}
