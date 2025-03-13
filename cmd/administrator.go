package cmd

import (
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/limanmys/limanctl/helpers"
	"github.com/limanmys/limanctl/models"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(administratorCmd)
}

func administrator() {
	db := helpers.DB

	password, _ := password.Generate(16, 8, 8, true, true)
	user := models.User{
		ID:            uuid.NewString(),
		Name:          "Administrator",
		Email:         "administrator@liman.dev",
		Password:      helpers.MakeHash(password),
		RememberToken: "",
		Status:        1,
		ForceChange:   true,
		AuthType:      "local",
		SessionTime:   -1,
		LastLoginAt:   time.Now(),
	}

	result := db.Create(&user)
	if result.RowsAffected < 1 {
		log.Fatal("Admin user already exists, if you lost password use <reset> command.")
	}

	color.Red("E-mail: \n" + "administrator@liman.dev" + "\n\n")
	color.Red("Password: \n" + password + "\n\n")
	color.Blue("When you login you will be prompted to change your password.")
}

var (
	administratorCmd = &cobra.Command{
		Use:   "administrator",
		Short: "Creates an administrator user on Liman",
		Long:  "Creates an administrator user on Liman",
		Run: func(cmd *cobra.Command, args []string) {
			administrator()
		},
	}
)
