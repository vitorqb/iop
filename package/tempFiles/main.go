package tempFiles

import (
	"io/ioutil"
	"os"
)

// Allows user to run a function with a temporary file path
type tempFile struct {}
func (tempFile) Run(f func(file *os.File) error) error {
	file, err := ioutil.TempFile("", "test-file")
	if err != nil {
		return err
	}
	err = f(file)
	if err != nil {
		return err
	}
	os.Remove(file.Name())
	return nil
}
func NewTempFile() tempFile {
	return tempFile{}
}

// Allows user to save a script to a temporary file, and run a script with it's path
type tempScript struct {
	body string
}

func (t tempScript) Run(f func(path string)) error {
	return NewTempFile().Run(func(file *os.File) error {
		_, err := file.WriteString(t.body)
		if err != nil {
			return err
		}
		err = file.Close()
		if err != nil {
			return err
		}
		err = os.Chmod(file.Name(), 0777)
		if err != nil {
			return err
		}
		f(file.Name())
		return err
	})
}

func NewTempScript(body string) tempScript {
	return tempScript{body: body}
}
