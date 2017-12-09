package storage

import (
	"github.com/while-loop/remember-me/remme/api/services/v1/record"
	"sync"
	"time"
)

type DataStore interface {
	AddLog(record *LogRecord) (*LogRecord, error)
	UpdateLog(record *LogRecord) (*LogRecord, error)
	GetLog(jobId uint64) (*LogRecord, error)
}

type LogRecord struct {
	Time  time.Time // timestamp
	JobID uint64    // uuid
	User  string    // remember-me user email

	tmu        sync.Mutex
	tries      uint64 // total amount of tries to change impl service pass
	TotalSites uint64 // total amount of sites user has

	fmu      sync.Mutex
	failures []*record.Failure // failures when logging in, changing pass, or unimpl host
}

func (lr *LogRecord) IncTries(amount uint64) {
	lr.tmu.Lock()
	defer lr.tmu.Unlock()
	lr.tries += amount
}

func (lr *LogRecord) AddFailure(hostname, email, reason, version string) {
	lr.fmu.Lock()
	defer lr.fmu.Unlock()

	lr.failures = append(lr.failures, &record.Failure{
		Hostname: hostname,
		Email:    email,
		Reason:   reason,
		Version:  version,
	})
}

func (lr *LogRecord) Failures() []*record.Failure {
	lr.fmu.Lock()
	defer lr.fmu.Unlock()
	return lr.failures
}
func (lr *LogRecord) Tries() uint64 {
	lr.tmu.Lock()
	defer lr.tmu.Unlock()
	return lr.tries
}
