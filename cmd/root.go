package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "limanctl",
		Short: "Liman MYS Control Application",
		Long:  `Limanctl is a CLI library for Liman MYS to track service statuses, maintaining system etc.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	out, err := exec.Command("id", "-u").Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			//fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
		} else {
			color.Red("failed to run id -u: %v", err)
			os.Exit(1)
		}
	}
	if strings.TrimSpace(string(out)) != "0" {
		color.Red("You must run Limanctl as sudo/root")
		os.Exit(1)
	}
}
