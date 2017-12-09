package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func init() {
	addCmd(versionCmd)
}

var versionCmd = cli.Command{
	Name: "version",
	Action: func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "%s version %s\n", c.App.Name, c.App.Version)
		return nil
	},
}
