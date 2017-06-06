package webservice

import (
	"fmt"
)

var (
	DEFAULT_SERVICES = []Webservice{NewFacebookWebservice()}
)

type Webservice interface {
	Login(email, password string) error
	Logout() error
	ChangePassword(email, oldpasswd, newpasswd string) error
	GetHostname() string
}

// Test cases
// Logged in successfully
// Logged in wrong passwd
// Logged in wrong email
// Logged in get login item to validate

// Change password wrong password
// Change passwd inv format
// change passws not matching
// change pass success
// change pass same pass
// change pass prev pass

func ChangeError(hostname, email, message string) error {
	return fmt.Errorf("Unable to change account %s password on hostname %s: %s", email, hostname, message)
}

func ParseError(hostname string) error {
	return fmt.Errorf("Unable to parse hostname: %s", hostname)
}

func ConnectError(hostname string) error {
	return fmt.Errorf("Unable to connect to web service: %s", hostname)
}

func AccountError(email string) error {
	return fmt.Errorf("Incorrect password for account %s", email)
}
