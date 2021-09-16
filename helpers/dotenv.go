package helpers

import (
	"github.com/joho/godotenv"
)

func GetKey(key string) string {
	dotenv, _ := godotenv.Read("/liman/server/.env")

	return dotenv[key]
}

func GetDbInfo() DB {
	return DB{
		Ip:       GetKey("DB_HOST"),
		Port:     GetKey("DB_PORT"),
		Username: GetKey("DB_USERNAME"),
		Password: GetKey("DB_PASSWORD"),
		Database: GetKey("DB_DATABASE"),
	}
}
