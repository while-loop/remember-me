package db

type DynamoDB struct {
}

func NewDynamoDB() *DynamoDB {
	return &DynamoDB{}
}

func (d *DynamoDB) AddLog(record *LogRecord) (*LogRecord, error) {
	return record, nil
}

func (d *DynamoDB) UpdateLog(record *LogRecord) (*LogRecord, error) {
	return record, nil
}

func (d *DynamoDB) GetServices() ([]string, error) {
	return nil, nil
}
