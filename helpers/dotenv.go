package helpers

import (
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func GetKey(key string) string {
	dotenv, err := godotenv.Read("/liman/server/.env")

	if err != nil {
		color.Red("You must run Limanctl as sudo/root")
		log.Fatal(err)
		os.Exit(1)
	}

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
