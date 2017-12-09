package webservice

import "fmt"

type ChangeError struct {
	Hostname string
	Email    string
	Message  string
}

func (ce ChangeError) Error() string {
	return fmt.Sprintf("Unable to change account %s password on hostname %s: %s", ce.Email, ce.Hostname, ce.Message)
}

type ParseError struct {
	Hostname string
}

func (pe ParseError) Error() string {
	return fmt.Sprintf("Unable to parse hostname: %s", pe.Hostname)
}

type ConnectError struct {
	Hostname string
}

func (ce ConnectError) Error() string {
	return fmt.Sprintf("Unable to connect to web service: %s", ce.Hostname)
}

type AccountError struct {
	Email    string
	Hostname string
}

func (ae AccountError) Error() string {
	return fmt.Sprintf("Incorrect password for account %s @ %s", ae.Email, ae.Hostname)
}
