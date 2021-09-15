package cmd

import (
	"fmt"
	"log"

	"github.com/limanmys/limanctl/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(dbStatus)
}

func checkDbStatus() {
	isAlive, err := helpers.CheckIfAlive(conn[0], conn[1], conn[2], conn[3], conn[4])

	if !isAlive {
		fmt.Println("PostgreSQL is working properly.")
	} else {
		fmt.Println("PostgreSQL is NOT working properly.")
		log.Fatal(err)
	}
}

var (
	conn = []string{
		helpers.GetKey("DB_HOST"),
		helpers.GetKey("DB_PORT"),
		helpers.GetKey("DB_USERNAME"),
		helpers.GetKey("DB_PASSWORD"),
		helpers.GetKey("DB_DATABASE"),
	}

	dbStatus = &cobra.Command{
		Use:   "check-database",
		Short: "Check if Liman MYS database is working",
		Long:  "Check if Liman MYS database is working",
		Run: func(cmd *cobra.Command, args []string) {
			checkDbStatus()
		},
	}
)
