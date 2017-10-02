package manager

import "fmt"

var (
	managers = map[string]RegisterFunc{}
)

type RegisterFunc func(email, password string) (Manager, error)

type Manager interface {
	GetEmail() string
	GetPassword(hostname, email string) (string, error)
	SavePassword(hostname, email, password string) error
	GetSites() []Site
}

func Register(name string, regFunc RegisterFunc) {
	managers[name] = regFunc
}

func GetManager(name, username, password string) (Manager, error) {
	val, exists := managers[name]
	if !exists {
		return nil, fmt.Errorf("manager does not exist: %s", name)
	}

	return val(username, password)
}

type Site struct {
	Hostname string
	Email    string
	Password string
}

func AccountDNE(hostname, email string) error {
	return fmt.Errorf("account %s does not exist on service: %s", email, hostname)
}
