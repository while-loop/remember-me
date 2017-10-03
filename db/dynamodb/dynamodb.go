package dynamodb

import (
	"github.com/while-loop/remember-me/db"
)

type DynamoDB struct {
}

func NewDynamoDB() *DynamoDB {
	return &DynamoDB{}
}

func (d *DynamoDB) AddLog(record *db.LogRecord) (*db.LogRecord, error) {
	return record, nil
}

func (d *DynamoDB) UpdateLog(record *db.LogRecord) (*db.LogRecord, error) {
	return record, nil
}

func (d *DynamoDB) GetLog(jobId uint64) (*db.LogRecord, error) {
	return nil, nil
}
