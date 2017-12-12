package stub

import (
	"github.com/while-loop/remember-me/remme/api/services/v1/record"
	"github.com/while-loop/remember-me/remme/log"
	"github.com/while-loop/remember-me/remme/storage"
)

type stubDB struct {
}

func New() storage.DataStore {
	return &stubDB{}
}

func (s *stubDB) AddLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	log.Info("AddLog", record)
	return record, nil
}
func (s *stubDB) UpdateLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	log.Info("UpdateLog", record)
	return record, nil
}
func (s *stubDB) GetLog(jobId uint64) (*storage.LogRecord, error) {
	log.Info("GetLog", jobId)
	return nil, nil
}

func (s *stubDB) AddEvent(event record.JobEvent) (record.JobEvent, error) {
	log.Info("AddEvent", event)
	return event, nil
}
