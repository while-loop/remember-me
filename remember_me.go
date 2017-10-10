package remme

import (
	"github.com/while-loop/remember-me/storage"
	"github.com/while-loop/remember-me/manager"
	"github.com/while-loop/remember-me/util"
	"github.com/while-loop/remember-me/webservice"
)

const (
	Version = "0.0.1"
	Release = "andromeda"
	Revision = 1
)

type PasswdFunc func() string

var (
	DefaultPasswdFunc = util.NewPasswordGen(32, true, true).Generate
)

func WebServices() map[string]webservice.Webservice {
	return webservice.Services()
}

func Hello() string {
    return "World"
}

func DefaultDB() storage.DataStore {
	return &storage.StubDB{}
}

func GetManager(name, username, password string) (manager.Manager, error) {
	return manager.GetManager(name, username, password)
}
