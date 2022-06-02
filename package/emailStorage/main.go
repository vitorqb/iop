package emailStorage

import "os"

type IEmailStorage interface {
	Put(email string) error
	Get() (string, error)
}

type fileEmailStorage struct {
	filePath string
}
// !!!! TODO Unify with tokenStorage
func(f fileEmailStorage) Put(email string) error {
	err := os.WriteFile(f.filePath, []byte(email), 0600)
	return err
}
// !!!! TODO Unify with tokenStorage
func(f fileEmailStorage) Get() (string, error) {
	token, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", nil
	}
	return string(token), nil
}

func New(filePath string) fileEmailStorage {
	return fileEmailStorage{ filePath: filePath }
}

// !!!! TODO Unify with tokenStorage
type inMemoryEmailStorage struct {
	Email string
}

// !!!! TODO Unify with tokenStorage
func (s *inMemoryEmailStorage) Put(email string) error {
	s.Email = email
	return nil
}
// !!!! TODO Unify with tokenStorage
func (s *inMemoryEmailStorage) Get() (string, error) {
	return s.Email, nil
}
// !!!! TODO Unify with tokenStorage
func NewInMemoryEmailStorage(initialEmail string) inMemoryEmailStorage {
	return inMemoryEmailStorage{Email: initialEmail}
}
