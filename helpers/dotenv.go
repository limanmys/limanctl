package helpers

import (
	"github.com/joho/godotenv"
)

func GetKey(key string) string {
	dotenv, _ := godotenv.Read("/liman/server/.env")

	return dotenv[key]
}
