package dynamodb

import (
	"github.com/while-loop/remember-me/remme/api/services/v1/record"
	"github.com/while-loop/remember-me/remme/storage"
)

type dynamoDB struct {
}

func New() storage.DataStore {
	return &dynamoDB{}
}

func (d *dynamoDB) AddEvent(event record.JobEvent) (record.JobEvent, error) {
	return event, nil
}

func (d *dynamoDB) AddLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	return record, nil
}

func (d *dynamoDB) UpdateLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	return record, nil
}

func (d *dynamoDB) GetLog(jobId uint64) (*storage.LogRecord, error) {
	return nil, nil
}
