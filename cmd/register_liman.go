package cmd

import (
	"log"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(registerLimanCmd)
}

func registerLiman() {
	// run php /liman/server/artisan register_liman
	_, err := exec.Command("/usr/bin/php", "/liman/server/artisan", "register_liman").Output()
	if err != nil {
		color.Red("Failed to register Liman IP to DB")
		log.Fatal(err)
	}
	color.Green("Liman IP has been registered successfully to DB")
}

var (
	registerLimanCmd = &cobra.Command{
		Use:   "register_liman",
		Short: "Registers current Liman to database for usage of HA features",
		Long:  "Registers current Liman to database for usage of HA features",
		Run: func(cmd *cobra.Command, args []string) {
			registerLiman()
		},
	}
)
