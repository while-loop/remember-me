package record

import (
	api "github.com/while-loop/remember-me/api/services/v1/record"
	"github.com/while-loop/remember-me/db"
)

func lr2Proto(lr *db.LogRecord) *api.LogRecord {
	return &api.LogRecord{
		JobId:      lr.JobID,
		Time:       uint64(lr.Time.Unix()),
		Email:      lr.User,
		Tries:      uint64(lr.Tries()),
		TotalSites: uint64(lr.TotalSites),
		Failures:   lr.Failures(),
	}
}
