package tokenStorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/tempFiles"
)

func Test_FileTokenStorage_PutAndGet(t *testing.T) {
	tempFiles.NewTempFile().Run(func(file *os.File) error {
		fileTokenStorage, err := New(file.Name())
		assert.Nil(t, err)

		getBeforePut, errBeforePut := fileTokenStorage.Get()
		assert.Equal(t, "", getBeforePut)
		assert.Nil(t, errBeforePut)

		fileTokenStorage.Put("FOO")
		getAfterPutFoo, errAfterPutFoo := fileTokenStorage.Get()
		assert.Equal(t, getAfterPutFoo, "FOO")
		assert.Nil(t, errAfterPutFoo)

		fileTokenStorage.Put("BAR")
		getAfterPutBar, errAfterPutBar := fileTokenStorage.Get()
		assert.Equal(t, getAfterPutBar, "BAR")
		assert.Nil(t, errAfterPutBar)

		return nil
	})
}