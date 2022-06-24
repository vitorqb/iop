package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_InMemorySimpleStorage_PutAndGet(t *testing.T) {
	inMemorySimpleStorage := NewInMemoryTokenStorage("FOO")

	read1, err := inMemorySimpleStorage.Get()
	assert.Nil(t, err)
	assert.Equal(t, "FOO", read1)

	err = inMemorySimpleStorage.Put("BAR")
	assert.Nil(t, err)

	read2, err := inMemorySimpleStorage.Get()
	assert.Nil(t, err)
	assert.Equal(t, "BAR", read2)
}
