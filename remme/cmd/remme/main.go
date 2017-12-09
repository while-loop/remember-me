package main

import (
	"github.com/urfave/cli"
	"os"
	"sync"
	"github.com/while-loop/remember-me/remme"
)

var app = struct {
	sync.Mutex
	a *cli.App
}{
	a: cli.NewApp(),
}

func main() {
	app.Lock()
	app.a.Name = "remember-me"
	app.a.Description = "Automatic password changer"
	app.a.Usage = app.a.Description
	app.a.Version = remme.Version
	app.a.HelpName = "remme"
	app.Unlock()

	app.a.Run(os.Args)
}

func addCmd(cmd cli.Command) {
	app.Lock()
	defer app.Unlock()

	app.a.Commands = append(app.a.Commands, cmd)
}
