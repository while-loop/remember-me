package storage

type StubDB struct {
}

func (s *StubDB) AddLog(record *LogRecord) (*LogRecord, error) {
	return record, nil
}
func (s *StubDB) UpdateLog(record *LogRecord) (*LogRecord, error) {
	return record, nil
}
func (s *StubDB) GetLog(jobId uint64) (*LogRecord, error) {
	return nil, nil
}
