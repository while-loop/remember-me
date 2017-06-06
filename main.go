package main

import (
	"./manager"
	"./webservice"
	"./db"
	"log"
)

const (
	MANAGER = "mem"
	VERSION = "1.0.0"
)

var (
	services = map[string]webservice.Webservice{}
	mngr     manager.Manager
)

func main() {

	username := "test"
	passwd := "test"

	var err error
	switch MANAGER {
	case "mem":
		mngr = genData()
	case "lastpass":
		mngr, err = manager.NewLastPassManager(username, passwd)
	default:
		mngr = genData()
	}

	if err != nil {
		log.Println(webservice.AccountError(username))
	}

	app := NewApp(db.NewDynamoDB(), webservice.DEFAULT_SERVICES...)
	app.ChangePasswords(nil, mngr)
}

func genPasswd() string {
	return "newpassword"
}

func genData() manager.Manager {
	mm := manager.NewMemManager()
	mm.SavePassword("facebook.com", "email", "password")
	return mm
}
