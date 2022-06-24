package storage

type InMemoryStorage struct {
	Value string
}
func (s *InMemoryStorage) Put(value string) error {
	s.Value = value
	return nil
}
func (s *InMemoryStorage) Get() (string, error) {
	return s.Value, nil
}
func NewInMemoryTokenStorage(initialValue string) InMemoryStorage {
	return InMemoryStorage{initialValue}
}
