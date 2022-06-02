package accountStorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/tempFiles"
)

func TestGetAndPut(t *testing.T) {
	err := tempFiles.NewTempFile().Run(func(file *os.File) error {
		storage := New(file.Name())

		gotten, err := storage.Get()
		if err != nil {
			return err
		}
		assert.Nil(t, err)
		assert.Equal(t, gotten, "")

		err = storage.Put("a@b.c")
		if err != nil {
			return err
		}

		gotten, err = storage.Get()
		if err != nil {
			return err
		}
		assert.Equal(t, gotten, "a@b.c")

		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
