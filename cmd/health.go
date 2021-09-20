package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(healthCmd)
}

var (
	healthCmd = &cobra.Command{
		Use:   "health",
		Short: "Run complete health check",
		Long:  "Run complete health check",
		Run: func(cmd *cobra.Command, args []string) {
			checkServices()
			fmt.Println(" \n --- \n ")
			checkDbStatus()
		},
	}
)
