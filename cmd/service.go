package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serviceCmd)
}

func checkServices() {
	success := true

	for _, service := range ServiceList {
		out, err := exec.Command("systemctl", "check", service).Output()
		if err != nil {
			if _, ok := err.(*exec.ExitError); ok {
				//fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
			} else {
				fmt.Printf("failed to run systemctl: %v", err)
				os.Exit(1)
			}
		}

		if strings.TrimSpace(string(out)) == "inactive" {
			success = false
			color.Red("%s status is %s", service, strings.TrimSpace(string(out)))
		} else {
			color.Green("%s status is %s", service, strings.TrimSpace(string(out)))
		}
	}

	if success {
		color.Green("\nAll services working as intended.")
	} else {
		color.Yellow("\nSome services is not working, Liman might not serve properly.")
	}
}

var (
	serviceCmd = &cobra.Command{
		Use:   "service",
		Short: "Checks health of Liman services",
		Long:  "Checks health of Liman services (render, webssh, system helper...)",
		Run: func(cmd *cobra.Command, args []string) {
			checkServices()
		},
	}

	ServiceList = []string{
		"liman-render",
		"liman-socket",
		"liman-system",
		"liman-vnc",
		"liman-webssh",
	}
)
