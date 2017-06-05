package webservice

import (
	"fmt"
)

type Webservice interface {
	Login(email, password string) error
	Logout() error
	ChangePassword(email, oldpasswd, newpasswd string) error
	GetDomain() string
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

func ChangeError(website, email, message string) error {
	return fmt.Errorf("Unable to change account %s password on website %s: %s", email, website, message)
}

func ParseError(website string) error {
	return fmt.Errorf("Unable to parse website: %s", website)
}

func ConnectError(website string) error {
	return fmt.Errorf("Unable to connect to web service: %s", website)
}

func AccountError(email string) error {
	return fmt.Errorf("Incorrect password for account %s", email)
}


