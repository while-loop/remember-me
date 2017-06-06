package manager

import "fmt"

type Manager interface {
	GetPassword(hostname, email string) (string, error)
	SavePassword(hostname, email, password string) error
	GetSites() []Site
}

// Tests
// login with correct password
// login with wrong password
// login with nonexistant email

// save password
// get password
// get all passwords

type Site struct {
	Hostname string
	Email string
	Password string
}

func AccountDNE(hostname, email string) error {
	return fmt.Errorf("Account %s does not exist on service: %s", email, hostname)
}