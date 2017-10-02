package main

import (
	"../../../remember-me"
	"../../../remember-me/db"
	"../../../remember-me/manager"
	"../../../remember-me/webservice"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func init() {
	addCmd(changeCmd)
}

var changeCmd = cli.Command{
	Name:    "change",
	Aliases: []string{"ch"},
	Flags: []cli.Flag{
		cli.StringFlag{Name: "m, manager", Value: "lastpass"},
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

		fmt.Printf("%v, %v", man.GetSites()[0].Hostname, man.GetSites()[0].Email)
		app := remember.NewApp(db.Default, webservice.Services()...)
		app.ChangePasswords(os.Stdout, man, remember.DefaultPasswdFunc)
		return nil
	},
}
