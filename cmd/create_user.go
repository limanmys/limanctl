package cmd

import (
	"log"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/google/uuid"
	"github.com/limanmys/limanctl/helpers"
	"github.com/limanmys/limanctl/models"
	"github.com/sethvargo/go-password/password"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createUserCmd)
}

func createUser(args []string) {
	if len(args) < 1 {
		log.Fatal("You didn't provide an e-mail address.")
	}
	if !strings.Contains(args[0], "@") {
		log.Fatal("Please provide a valid e-mail address.")
	}

	db := helpers.DB

	var user models.User
	if result := db.First(&user, "email = ?", args[0]); result.Error == nil {
		log.Fatal("User found on Liman, provide different e-mail address.")
	}

	password, _ := password.Generate(16, 8, 8, true, true)
	datetime := time.Now().Format(time.RFC3339)

	user = models.User{
		ID:          uuid.New().String(),
		Name:        "New User",
		Email:       args[0],
		Password:    helpers.MakeHash(password),
		CreatedAt:   datetime,
		UpdatedAt:   datetime,
		LastLoginAt: datetime,
		LastLoginIP: "0.0.0.0",
		ForceChange: true,
		Status:      1,
		AuthType:    "local",
	}

	result := db.Create(&user)
	if result.RowsAffected < 1 {
		log.Fatal("An error occured while creating user.")
	}

	color.Red("E-mail: \n" + args[0] + "\n\n")
	color.Red("New password: \n" + password + "\n\n")
	color.Blue("When you login you will be prompted to change your password.")
}

var (
	createUserCmd = &cobra.Command{
		Use:   "create-user",
		Short: "Create new Liman user, needs [email] arg",
		Long:  "Create new Liman user, needs [email] arg",
		Run: func(cmd *cobra.Command, args []string) {
			createUser(args)
		},
	}
)
