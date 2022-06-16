package testUtils

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
	"testing"

	"github.com/vitorqb/iop/package/tempFiles"
)

func GetTestDataFilePath(testDataFileName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "/test_data", testDataFileName), nil
}

func GetTestDataDirectory() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(path, "/test_data"), nil	
}

func RenderTemplateTestFile(
	t *testing.T,
	templateFileName string,
	data interface{},
) string {
	tmplFile, err := GetTestDataFilePath(templateFileName)
	if err != nil {
		t.Error(err)
	}
	tmpl, err := template.New(templateFileName).ParseFiles(tmplFile)
	if err != nil {
		t.Error(err)
	}
	var output bytes.Buffer
	err = tmpl.Execute(&output, data)
	if err != nil {
		t.Error(err)
	}
	return tempFiles.NewTestTempScript(t, output.String())
}
