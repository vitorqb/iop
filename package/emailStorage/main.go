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
