package storage

import "os"

// A simple file-based implementation of ISimpleStorage
type FileSimpleStorage struct {
	filePath string
}

func (f FileSimpleStorage) Put(value string) error {
	err := os.WriteFile(f.filePath, []byte(value), 0600)
	if err != nil {
		return err
	}
	return nil
}

func (f FileSimpleStorage) Get() (string, error) {
	token, err := os.ReadFile(f.filePath)
	if err != nil {
		return "", nil
	}
	return string(token), nil
}

func NewFileSimpleStorage(filePath string) FileSimpleStorage {
	return FileSimpleStorage {
		filePath: filePath,
	}
}
