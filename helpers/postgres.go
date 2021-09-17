package helpers

import (
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB = connectGorm()

func ConnectionQuery() string {
	conn := GetDbInfo()

	port, _ := strconv.Atoi(conn[1])
	connectionQuery := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", conn[0], port, conn[2], conn[3], conn[4])

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
	db, _ := DB.DB()

	err := db.Ping()
	info := GetDbInfo()
	if err == nil {
		fmt.Printf("Connected successfully to %s:%s\n", info[0], info[1])
	}

	return err != nil, err
}
