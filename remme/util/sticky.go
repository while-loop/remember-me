package util

type StickyError struct {
	err error
}

func (s *StickyError) Swallow(err error) {
	if s.err == nil {
		s.err = err
	}
}

func (s *StickyError) Error() error {
	return s.err
}

func (s StickyError) HasError() bool {
	return s.err != nil
}
