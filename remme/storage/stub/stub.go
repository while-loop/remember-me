package stub

import (
	"github.com/while-loop/remember-me/remme/storage"
)

type StubDB struct {
}

func New() storage.DataStore{
	return &StubDB{}
}

func (s *StubDB) AddLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	return record, nil
}
func (s *StubDB) UpdateLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	return record, nil
}
func (s *StubDB) GetLog(jobId uint64) (*storage.LogRecord, error) {
	return nil, nil
}
