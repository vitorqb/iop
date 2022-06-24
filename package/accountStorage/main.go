package accountStorage

import "github.com/vitorqb/iop/package/storage"

// TODO Ensure we create parents if needed!
func New(filePath string) storage.ISimpleStorage {
	return storage.NewFileSimpleStorage(filePath)
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
