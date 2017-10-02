package main

import (
	"../../../remember-me"
	"github.com/urfave/cli"
	"os"
)

var (
	app = cli.NewApp()
)

func main() {
	app.Name = "remember-me"
	app.Description = "Automatic password changer"
	app.Usage = app.Description
	app.Version = remember.Version
	app.HelpName = "remme"
	app.Run(os.Args)
}

func addCmd(cmd cli.Command) {
	app.Commands = append(app.Commands, cmd)
}
