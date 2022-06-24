package tokenStorage

import (
	"os"
	"path/filepath"

	"github.com/vitorqb/iop/package/storage"
)

// Helper functions for default implementations
const DEFAULT_TOKEN_FILE_RELATIVE_TO_USER_DIR = ".iop/credentials"

func getDefaultTokenFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	tokenDir := filepath.Join(homeDir, DEFAULT_TOKEN_FILE_RELATIVE_TO_USER_DIR)
	return tokenDir, nil
}
func createParents(filePath string) error {
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	return err
}

func New(filePath string) (storage.ISimpleStorage , error) {
	if filePath == "" {
		filePath1, err := getDefaultTokenFile()
		if err != nil {
			return nil, err
		}
		filePath = filePath1
	}
	err := createParents(filePath)
	if err != nil {
		return nil, err
	}
	return storage.NewFileSimpleStorage(filePath), nil
}

// An implementation that stores the token in memory (usefull 4 tests)
type inMemoryTokenStorage struct {
	Token string
}

func (s *inMemoryTokenStorage) Put(token string) error {
	s.Token = token
	return nil
}
func (s *inMemoryTokenStorage) Get() (string, error) {
	return s.Token, nil
}
func NewInMemoryTokenStorage(initialToken string) inMemoryTokenStorage {
	return inMemoryTokenStorage{Token: initialToken}
}
