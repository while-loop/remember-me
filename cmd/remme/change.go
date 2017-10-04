package main

import (
	"fmt"
	"github.com/urfave/cli"
	"github.com/while-loop/remember-me"
	"github.com/while-loop/remember-me/api/services/v1/changer"
	"strings"
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
		man, err := remme.GetManager(manStr, email, password)
		if err != nil {
			fmt.Fprint(c.App.ErrWriter, err)
			return err
		}

		app := remme.NewApp(remme.DefaultDB(), remme.WebServices())
		statusChan := make(chan changer.Status)
		go app.ChangePasswords(statusChan, man, remme.DefaultPasswdFunc)

		for status := range statusChan {
			fmt.Fprintln(c.App.Writer, status)
		}
		return nil
	},
}
