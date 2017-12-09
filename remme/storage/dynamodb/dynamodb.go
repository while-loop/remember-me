package dynamodb

import (
	"github.com/while-loop/remember-me/remme/storagetorage"
)

type DynamoDB struct {
}

func New() storage.DataStore {
	return &DynamoDB{}
}

func (d *DynamoDB) AddLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	return record, nil
}

func (d *DynamoDB) UpdateLog(record *storage.LogRecord) (*storage.LogRecord, error) {
	return record, nil
}

func (d *DynamoDB) GetLog(jobId uint64) (*storage.LogRecord, error) {
	return nil, nil
}
