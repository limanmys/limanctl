package cmd

import (
	"log"

	"github.com/fatih/color"
	"github.com/limanmys/limanctl/helpers"
	"github.com/limanmys/limanctl/models"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resetPwCmd)
}

func resetPassword(args []string) {
	if len(args) < 1 {
		log.Fatal("Please provide e-mail address.")
	}

	db := helpers.DB

	var user models.User
	if result := db.First(&user, "email = ?", args[0]); result.Error != nil {
		log.Fatal("User not found on Liman")
	}

	password, _ := password.Generate(16, 8, 8, true, true)

	user.Password = helpers.MakeHash(password)
	user.ForceChange = true

	result := db.Save(&user)
	if result.RowsAffected < 1 {
		log.Fatal("Could not change password.")
	}

	color.Red("E-mail: \n" + args[0] + "\n\n")
	color.Red("New password: \n" + password + "\n\n")
	color.Blue("When you login you will be prompted to change your password.")
}

var (
	resetPwCmd = &cobra.Command{
		Use:   "reset",
		Short: "Reset a Liman user's password, requires [email] argument",
		Long:  "Reset a Liman user's password, requires [email] argument, example usage: sudo limanctl reset-password administrator@liman.dev",
		Run: func(cmd *cobra.Command, args []string) {
			resetPassword(args)
		},
	}
)
