package accountStorage

import "os"

type IAccountStorage interface {
	Put(account string) error
	Get() (string, error)
}

type fileAccountStorage struct {
	filePath string
}
// !!!! TODO Unify with tokenStorage
func(f fileAccountStorage) Put(account string) error {
	err := os.WriteFile(f.filePath, []byte(account), 0600)
	return err
}
// !!!! TODO Unify with tokenStorage
func(f fileAccountStorage) Get() (string, error) {
	token, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", nil
	}
	return string(token), nil
}

func New(filePath string) fileAccountStorage {
	return fileAccountStorage{ filePath: filePath }
}

// !!!! TODO Unify with tokenStorage
type inMemoryAccountStorage struct {
	Account string
}

// !!!! TODO Unify with tokenStorage
func (s *inMemoryAccountStorage) Put(account string) error {
	s.Account = account
	return nil
}
// !!!! TODO Unify with tokenStorage
func (s *inMemoryAccountStorage) Get() (string, error) {
	return s.Account, nil
}
// !!!! TODO Unify with tokenStorage
func NewInMemoryAccountStorage(initialAccount string) inMemoryAccountStorage {
	return inMemoryAccountStorage{Account: initialAccount}
}
