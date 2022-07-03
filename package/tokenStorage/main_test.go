package tokenStorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/pmwrap/package/tempFiles"
)

func Test_FileTokenStorage_PutAndGet(t *testing.T) {
	err := tempFiles.NewTempFile().Run(func(file *os.File) error {
		fileTokenStorage, err := New(file.Name())
		assert.Nil(t, err)

		getBeforePut, errBeforePut := fileTokenStorage.Get()
		assert.Equal(t, "", getBeforePut)
		assert.Nil(t, errBeforePut)

		err = fileTokenStorage.Put("FOO")
		if err != nil {
			t.Error(err)
		}
		getAfterPutFoo, errAfterPutFoo := fileTokenStorage.Get()
		assert.Equal(t, getAfterPutFoo, "FOO")
		assert.Nil(t, errAfterPutFoo)

		err = fileTokenStorage.Put("BAR")
		if err != nil {
			t.Error(err)
		}
		getAfterPutBar, errAfterPutBar := fileTokenStorage.Get()
		assert.Equal(t, getAfterPutBar, "BAR")
		assert.Nil(t, errAfterPutBar)

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
