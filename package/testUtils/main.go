package testUtils

import (
	"os"
	"path/filepath"
)

func GetTestDataFilePath(testDataFileName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "/test_data", testDataFileName), nil
}
