package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/limanmys/limanctl/helpers"
	"github.com/limanmys/limanctl/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCertificateCmd)
	rootCmd.AddCommand(fixCertificateCmd)
}

func checkCertificate() {
	db := helpers.DB

	var certificates []models.Certificate
	if result := db.Find(&certificates); result.Error != nil {
		log.Fatal("No certificate exists on Liman")
	}

	flag := true
	fix := false
	for _, certs := range certificates {
		certFile := fmt.Sprintf("/usr/local/share/ca-certificates/liman-%s_%s.crt", certs.ServerHostname, certs.Origin)
		certContent, err := helpers.GetFileContents(certFile)
		if err != nil {
			log.Println("Unable to open certificate file " + certFile + "\n" + err.Error())
		}

		remoteCert, connErr := helpers.GetCertificate(fmt.Sprintf("%s:%s", certs.ServerHostname, certs.Origin))
		if connErr != nil {
			log.Println("Unable to connect to " + certs.ServerHostname)
		}

		if strings.Compare(remoteCert, certContent) == 0 {
			color.Green(fmt.Sprintf("%s:%s certificate is matching with the Liman's system", certs.ServerHostname, certs.Origin))
		} else {
			color.Red(fmt.Sprintf("%s:%s certificate is not matching with the Liman's system", certs.ServerHostname, certs.Origin))
			flag = false

			color.Cyan("Would you like to fix this certificate issue? [y/N]")
			choice := ""
			fmt.Scanln(&choice)

			if strings.Contains(strings.ToLower(choice), "y") {
				fix = true
				fixCertificate(fmt.Sprintf("%s:%s", certs.ServerHostname, certs.Origin))
			}
		}
	}

	if flag {
		color.Green("All certificates are valid")
	} else {
		color.Red("One (or more) of the certificates are invalid")

		if fix {
			color.Cyan("Looks like you fixed an issue, would you like to run certificate check again? [y/N]")
			choice := ""
			fmt.Scanln(&choice)

			if strings.Contains(strings.ToLower(choice), "y") {
				checkCertificate()
			}
		}
	}
}

func fixCertificate(conn string) {
	remoteCert, connErr := helpers.GetCertificate(conn)
	if connErr != nil {
		log.Println("Unable to connect to " + conn)
	}

	f, err := os.Create("/usr/local/share/ca-certificates/liman-" + strings.Replace(conn, ":", "_", -1) + ".crt")
	if err != nil {
		log.Println("Unable to create certificate file.")
	}
	defer f.Close()

	_, writeErr := f.WriteString(remoteCert)
	if writeErr != nil {
		log.Println("An error occured while writing the certificate into the file.")
	}

	updateErr := exec.Command("update-ca-certificates").Run()
	if updateErr != nil {
		log.Println("An error occured while running update-ca-certificates")
	}
}

var (
	checkCertificateCmd = &cobra.Command{
		Use:   "check-certificates",
		Short: "Check if certificates in Liman is valid",
		Long:  "Check if certificates in Liman is valid",
		Run: func(cmd *cobra.Command, args []string) {
			checkCertificate()
		},
	}

	fixCertificateCmd = &cobra.Command{
		Use:   "fix-certificate",
		Short: "Fix the certificate for the specified connection, needs ip:port as arg",
		Long:  "Fix the certificate for the specified connection, needs ip:port as arg",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				return
			}
			fixCertificate(args[0])
		},
	}
)
