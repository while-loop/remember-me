package main

import (
	"../../remember-me"
	"../db"
	"../manager"
	"../webservice"
	"fmt"
	"os"
)

func main() {
	man, err := manager.GetManager("", "@gmail.com", "")
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}
	app := remember.NewApp(db.Default, webservice.Services()...)
	app.ChangePasswords(os.Stdout, man, remember.DefaultPasswdFunc)
}
