package tokenStorage

import (
	"os"
	"path/filepath"
)

type ITokenStorage interface {
	Put(token string) error
	Get() (string, error)
}

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
	dir := filepath.Dir(filePath)
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fileTokenStorage{}, err
	}
	return fileTokenStorage{filePath: filePath}, nil
}
