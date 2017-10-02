package remember

import (
	"./db"
	"./manager"
	"./webservice"
	"log"
)

const (
	MANAGER = "mem"
)

func main() {
	var mngr manager.Manager
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
		log.Println(webservice.AccountError{username, MANAGER})
	}

	app := NewApp(db.NewDynamoDB(), webservice.Services()...)
	app.ChangePasswords(nil, mngr, DefaultPasswdFunc)
}

func genData() manager.Manager {
	mm := manager.NewMemManager()
	mm.SavePassword("facebook.com", "email", "password")
	return mm
}
