package tokenStorage

import (
	"os"
	"path/filepath"
)

// Helper functions for default implementations
const DEFAULT_TOKEN_FILE_RELATIVE_TO_USER_DIR = ".pmwrap/credentials"

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

// An interface capable of putting and getting a token
type ITokenStorage interface {
	Put(token string) error
	Get() (string, error)
}

// An implementation that stores the token in a file
type fileTokenStorage struct {
	filePath string
}

func (f fileTokenStorage) Put(token string) error {
	err := os.WriteFile(f.filePath, []byte(token), 0600)
	if err != nil {
		return err
	}
	return nil
}
func (f fileTokenStorage) Get() (string, error) {
	token, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", nil
	}
	return string(token), nil
}
func New(filePath string) (fileTokenStorage, error) {
	if filePath == "" {
		filePath1, err := getDefaultTokenFile()
		if err != nil {
			return fileTokenStorage{}, err
		}
		filePath = filePath1
	}
	err := createParents(filePath)
	if err != nil {
		return fileTokenStorage{}, err
	}
	return fileTokenStorage{filePath: filePath}, nil
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
