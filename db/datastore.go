package db

import (
	"time"
	"sync"
)

type DataStore interface {
	AddLog(record *LogRecord) (*LogRecord, error)
	UpdateLog(record *LogRecord) (*LogRecord, error)
	GetServices() ([]string, error)
}

var (
	Default = NewDynamoDB()
)

type LogRecord struct {
	Time  time.Time // timestamp
	JobID uint64    // uuid
	User  string    // remember-me user email

	tmu        sync.Mutex
	tries      uint // total amount of tries to change impl service pass
	TotalSites uint // total amount of sites user has

	fmu      sync.Mutex
	failures []Failure // failures when logging in, changing pass, or unimpl host
}

type Failure struct {
	Hostname string
	Email    string
	Reason   string
	Version  string
}

func (lr *LogRecord) IncTries(amount uint) {
	lr.tmu.Lock()
	defer lr.tmu.Unlock()
	lr.tries += amount
}

func (lr *LogRecord) AddFailure(hostname, email, reason, version string) {
	lr.fmu.Lock()
	defer lr.fmu.Unlock()

	lr.failures = append(lr.failures, Failure{
		Hostname: hostname,
		Email:    email,
		Reason:   reason,
		Version:  version,
	})
}

func (lr *LogRecord) Failures() []Failure {
	lr.fmu.Lock()
	defer lr.fmu.Unlock()
	return lr.failures
}
func (lr *LogRecord) Tries() uint {
	lr.tmu.Lock()
	defer lr.tmu.Unlock()
	return lr.tries
}
