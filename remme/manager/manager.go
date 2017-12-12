package manager

import (
	"fmt"
	"strings"
	"sync"
)

var managers = struct {
	sync.Mutex
	m map[string]RegisterFunc
}{
	m: map[string]RegisterFunc{},
}

type RegisterFunc func(email, password string) (Manager, error)

type Manager interface {
	GetEmail() string
	GetPassword(hostname, email string) (string, error)
	SavePassword(hostname, email, password string) error
	GetSites() ([]Site, error)
	Name() string
}

func Register(name string, regFunc RegisterFunc) {
	managers.Lock()
	defer managers.Unlock()

	managers.m[name] = regFunc
}

func GetManager(name, username, password string) (Manager, error) {
	managers.Lock()
	defer managers.Unlock()

	name = strings.ToLower(name)
	val, exists := managers.m[name]
	if !exists {
		return nil, fmt.Errorf("password manager %s is unavailable", name)
	}

	return val(username, password)
}

func GetManagers() []string {
	managers.Lock()
	defer managers.Unlock()

	mans := []string{}
	for man := range managers.m {
		mans = append(mans, man)
	}
	return mans
}

type Site struct {
	Hostname string
	Email    string
	Password string
}

func AccountDNE(hostname, email string) error {
	return fmt.Errorf("account %s does not exist on service: %s", email, hostname)
}
