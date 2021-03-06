package main

import (
	"fmt"
	"strings"

	"github.com/urfave/cli"
	"github.com/while-loop/remember-me/remme"
	"github.com/while-loop/remember-me/remme/log"
	"github.com/while-loop/remember-me/remme/manager"
	"github.com/while-loop/remember-me/remme/storage/stub"
	"github.com/while-loop/remember-me/remme/util"
	"github.com/while-loop/remember-me/remme/webservice"
)

func init() {
	addCmd(changeCmd)
}

var changeCmd = cli.Command{
	Name:    "change",
	Aliases: []string{"ch"},
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "m, manager",
			Value: "lastpass",
			Usage: "which account manager to use"},
	},
	Usage:     "change passwords for a given manager",
	ArgsUsage: "[-m manager] <email> <password>",
	Action: func(c *cli.Context) error {
		if c.NArg() != 2 {
			cli.ShowCommandHelp(c, "change")
			return nil
		}

		manStr, email, password := strings.ToLower(c.String("m")), c.Args().Get(0), c.Args().Get(1)
		man, err := manager.GetManager(manStr, email, password)
		if err != nil {
			fmt.Fprintln(c.App.ErrWriter, err)
			return err
		}

		app := remme.NewApp(stub.New(), webservice.Services())
		jobId := app.ChangePasswords(man, util.DefaultPasswdFunc)
		log.Info("Job ID", jobId)
		return nil
	},
}
