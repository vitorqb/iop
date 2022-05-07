package testUtils

import (
	"path/filepath"
	"os"
)

func GetTestDataFilePath(testDataFileName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "/test_data", testDataFileName), nil
}
