package manager

import "fmt"

type Manager interface {
	GetPassword(hostname, email string) (string, error)
	SavePassword(hostname, email, password string) error
	GetSites() []Site
}

type Site struct {
	Domain string
	Email string
	Password string
}


func AccountDNE(host, email string) error {
	return fmt.Errorf("Account %s does not exist on service: %s", email, host)
}