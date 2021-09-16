package cmd

import (
	"fmt"
	"log"

	"github.com/limanmys/limanctl/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

func printVersion() {
	version, err := helpers.GetFileContents("/liman/server/storage/VERSION")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Liman MYS Version: %s - ", version)
	ver_code, err := helpers.GetFileContents("/liman/server/storage/VERSION_CODE")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ver_code)
}

var (
	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Version of Liman",
		Long:  "LimanMYS version number",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	}
)
