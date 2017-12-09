package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/while-loop/remember-me/remme/api/services/v1/changer"
	"strings"
	"github.com/while-loop/remember-me/remme/manager"
	"github.com/while-loop/remember-me/remme/storage/stub"
	"github.com/while-loop/remember-me/remme/webservice"
	"github.com/while-loop/remember-me/remme/util"
	"github.com/while-loop/remember-me/remme"
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
			fmt.Fprint(c.App.ErrWriter, err)
			return err
		}

		app := remme.NewApp(stub.New(), webservice.Services())
		statusChan := make(chan changer.Status)
		go app.ChangePasswords(statusChan, man, util.DefaultPasswdFunc)

		for status := range statusChan {
			fmt.Fprintln(c.App.Writer, status)
		}
		return nil
	},
}
