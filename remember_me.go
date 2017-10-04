package remme

import (
	"github.com/while-loop/remember-me/db"
	"github.com/while-loop/remember-me/managers"
	"github.com/while-loop/remember-me/util"
	"github.com/while-loop/remember-me/webservices"
)

const (
	Version = "0.0.1"
	Release = "andromeda"
)

type PasswdFunc func() string

var (
	DefaultPasswdFunc = util.NewPasswordGen(32, true, true).Generate
)

func WebServices() map[string]webservices.Webservice {
	return webservices.Services()
}

func DefaultDB() db.DataStore {
	return &db.StubDB{}
}

func GetManager(name, username, password string) (managers.Manager, error) {
	return managers.GetManager(name, username, password)
}
