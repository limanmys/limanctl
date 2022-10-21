package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/limanmys/limanctl/helpers"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(backup)
}

func backupDb() {
	path := ""
	if _, err := os.Stat("/usr/bin/pg_dump"); err == nil {
		// debian
		path = "/usr/bin/pg_dump"
	} else if errors.Is(err, os.ErrNotExist) {
		// redhat
		path = "/usr/pgsql-15/bin/pg_dump"
	}

	err := os.MkdirAll("/liman/sql-backups", 0770)
	if err != nil {
		color.Red("Cannot create backup")
	}

	limanuser, err := user.Lookup("liman")
	if err != nil {
		log.Fatal(err)
	}
	limanuid, _ := strconv.Atoi(limanuser.Uid)
	limangid, _ := strconv.Atoi(limanuser.Gid)

	err = os.Chown("/liman/sql-backups", limanuid, limangid)
	if err != nil {
		log.Fatal(err)
	}

	db := helpers.GetDbInfo()

	currentTime := time.Now()

	fileName := fmt.Sprintf("/liman/sql-backups/liman-backup-%s.sql", currentTime.Format("2006-01-02-15-04-05"))

	out := exec.Command("sh", "-c", fmt.Sprintf(
		"%s --dbname=postgresql://%s:%s@%s:%s/%s > %s",
		path,
		db[2],
		db[3],
		db[0],
		db[1],
		db[4],
		fileName,
	)).Run()

	if out == nil {
		color.Green("Backup generated under " + fileName)
	} else {
		color.Red("Error occured while generating the backup.")
		log.Fatal(out)
	}
}

var (
	backup = &cobra.Command{
		Use:   "backupdb",
		Short: "Generate a backup under /liman/sql-backups",
		Long:  "Generate a backup under /liman/sql-backups",
		Run: func(cmd *cobra.Command, args []string) {
			backupDb()
		},
	}
)
