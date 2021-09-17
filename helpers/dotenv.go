package helpers

import (
	"github.com/joho/godotenv"
)

func GetKey(key string) string {
	dotenv, _ := godotenv.Read("/liman/server/.env")

	return dotenv[key]
}

func GetDbInfo() []string {
	return []string{
		GetKey("DB_HOST"),
		GetKey("DB_PORT"),
		GetKey("DB_USERNAME"),
		GetKey("DB_PASSWORD"),
		GetKey("DB_DATABASE"),
	}
}
