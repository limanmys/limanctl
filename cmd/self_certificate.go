package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(certificateCmd)
}

func createCertificate(args []string) {
	if len(args) < 1 {
		log.Fatal("You didn't provide an e-mail address.")
	}

	color.HiBlue("Certificate generating...")
	_, err := exec.Command("openssl", "req", "-x509", "-newkey", "rsa:4096", "-sha256", "-days", "3650", "-nodes", "-keyout", "/liman/certs/liman.key", "-out", "/liman/certs/liman.crt", "-subj", "/CN="+args[0]).Output()
	if err != nil {
		log.Fatal(err)
	}
	color.Green("Certificate generated!\n\n")

	color.HiBlue("Fixing permissions...")
	limanuser, err := user.Lookup("liman")
	if err != nil {
		log.Fatal(err)
	}
	limanuid, _ := strconv.Atoi(limanuser.Uid)
	limangid, _ := strconv.Atoi(limanuser.Gid)

	err = os.Chown("/liman/certs/liman.key", limanuid, limangid)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Chown("/liman/certs/liman.crt", limanuid, limangid)
	if err != nil {
		log.Fatal(err)
	}
	color.Green("Fixed permissions!\n\n")

	color.HiBlue("Restarting Liman Web Service...")
	out, err := exec.Command("systemctl", "restart", "nginx").Output()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			//fmt.Printf("systemctl finished with non-zero: %v\n", exitErr)
		} else {
			fmt.Printf("failed to run systemctl: %v", err)
			os.Exit(1)
		}
	}
	fmt.Printf("%s", string(out))
	color.Green("Liman Web Service restarted!\n\n")

	color.Magenta("Self signed certificate installed! To use it you should re-open your browser.")
}

var (
	certificateCmd = &cobra.Command{
		Use:   "cert",
		Short: "Generate self signed certificate for Liman Core, needs CN name as argument.",
		Long:  "Generate self signed certificate for Liman Core, needs CN name as argument.",
		Run: func(cmd *cobra.Command, args []string) {
			createCertificate(args)
		},
	}
)
