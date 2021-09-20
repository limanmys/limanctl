package cmd

import (
	"log"
	"os"

	"github.com/limanmys/limanctl/helpers"
	"github.com/limanmys/limanctl/models"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(extensionCmd)
}

func listExtensions() {
	db := helpers.DB

	var extensions []models.Extension
	if result := db.Find(&extensions); result.Error != nil {
		log.Fatal("No extension found on Liman")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Version"})

	for _, v := range extensions {
		table.Append([]string{v.Name, v.Version})
	}

	table.Render()
}

var (
	extensionCmd = &cobra.Command{
		Use:   "extensions",
		Short: "List all installed extensions",
		Long:  "List all installed extensions",
		Run: func(cmd *cobra.Command, args []string) {
			listExtensions()
		},
	}
)
