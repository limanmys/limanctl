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
	rootCmd.AddCommand(listUsersCmd)
}

func listUsers() {
	db := helpers.DB

	var users []models.User
	if result := db.Find(&users); result.Error != nil {
		log.Fatal("User not found on Liman")
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"User Type", "Name", "E-mail", "Last Login IP", "Auth Type"})

	for _, v := range users {
		var userType string
		if v.Status == 1 {
			userType = "Admin"
		} else {
			userType = "User"
		}
		table.Append([]string{
			userType, v.Name, v.Email, v.LastLoginIP.String, v.AuthType,
		})
	}

	table.Render()
}

var (
	listUsersCmd = &cobra.Command{
		Use:   "users",
		Short: "List Liman users",
		Long:  "You can list all Liman users",
		Run: func(cmd *cobra.Command, args []string) {
			listUsers()
		},
	}
)
