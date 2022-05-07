package tempScript

import (
	"io/ioutil"
	"os"
)

type tempScript struct {
	body string
}

func New(body string) tempScript {
	return tempScript{body: body}
}

func (t tempScript) Run(f func(path string)) error {
	file, err := ioutil.TempFile("", "test-file")
	if err != nil {
		return err
	}
	_, err = file.WriteString(t.body)
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
	err = os.Remove(file.Name())
	return err
}
