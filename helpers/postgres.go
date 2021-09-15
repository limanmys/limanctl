package helpers

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func CheckIfAlive(ip string, port string, username string, password string, database string) (bool, error) {
	port_, _ := strconv.Atoi(port)
	connectionQuery := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", ip, port_, username, password, database)

	db, err := sql.Open("postgres", connectionQuery)
	if err != nil {
		return false, err
	}

	defer db.Close()

	err = db.Ping()
	if err == nil {
		fmt.Printf("Connected successfully to %s:%s\n", ip, port)
	}

	return err != nil, err
}
