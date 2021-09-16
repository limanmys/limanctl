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
	isAlive, err := helpers.CheckIfAlive()

	if !isAlive {
		fmt.Println("PostgreSQL is working properly.")
	} else {
		fmt.Println("PostgreSQL is NOT working properly.")
		log.Fatal(err)
	}
}

var (
	dbStatus = &cobra.Command{
		Use:   "database",
		Short: "Check if Liman MYS database is working",
		Long:  "Check if Liman MYS database is working",
		Run: func(cmd *cobra.Command, args []string) {
			checkDbStatus()
		},
	}
)
