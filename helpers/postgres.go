package helpers

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Ip       string
	Port     string
	Username string
	Password string
	Database string
}

var Db = connectGorm()

func ConnectionQuery() string {
	conn := GetDbInfo()

	port, _ := strconv.Atoi(conn.Port)
	connectionQuery := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conn.Ip, port, conn.Username, conn.Password, conn.Database)

	return connectionQuery
}

func connectGorm() *gorm.DB {
	gormDB, err := gorm.Open(postgres.Open(ConnectionQuery()), &gorm.Config{
		// Log bastırmak istemediğimizde bu kısmı açacaz
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal(err)
	}

	return gormDB
}

func CheckIfAlive() (bool, error) {
	db, _ := Db.DB()

	err := db.Ping()
	if err == nil {
		fmt.Printf("Connected successfully to %s:%s\n", GetDbInfo().Ip, GetDbInfo().Port)
	}

	return err != nil, err
}
