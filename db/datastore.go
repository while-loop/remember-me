package db

import "time"

type DataStore interface {
	AddLog(record LogRecord) (LogRecord, error)
	UpdateLog(record LogRecord) (LogRecord, error)
	GetServices() ([]string, error)
}

type LogRecord struct {
	Time       time.Time	// timestamp
	JobID      uint64		// uuid
	User       string		// remember-me user email
	Fails      uint			// total amount of fails trying to change impl service pass
	Tries      uint			// total amount of tries to change impl service pass
	TotalSites uint			// total amount of sites user has
	Failures   []Failure	// failures when logging in, changing pass, or unimpl host
}

type Failure struct {
	Hostname string
	Email    string
	Reason   string
	Version  string
}

func (lr *LogRecord) AddFailure(hostname, email, reason, version string) {
	lr.Failures = append(lr.Failures, Failure{
		Hostname: hostname,
		Email:    email,
		Reason:   reason,
		Version:  version,
	})
}
