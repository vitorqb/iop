package storage

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FileSimpleStorage_PutAndGet(t *testing.T) {
	tempFile, err := ioutil.TempFile("", "")
	assert.Nil(t, err)
	fileSimpleStorage := NewFileSimpleStorage(tempFile.Name())

	read1, err := fileSimpleStorage.Get()
	assert.Nil(t, err)
	assert.Equal(t, "", read1)

	err = fileSimpleStorage.Put("FOO")
	assert.Nil(t, err)

	read2, err := fileSimpleStorage.Get()
	assert.Nil(t, err)
	assert.Equal(t, "FOO", read2)
}
