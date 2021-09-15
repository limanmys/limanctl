package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version of Liman",
		Long:  "LimanMYS version number",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version 0.1")
		},
	}
)
