package commandRunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandRunnerCallsCommand(t *testing.T) {
	commandRunner := CommandRunner{}
	res, err := commandRunner.Run("echo", "-n", "FOO")
	assert.Nil(t, err)
	assert.Equal(t, "FOO", string(res))
}

func TestCommandRunnerError(t *testing.T) {
	commandRunner := CommandRunner{}
	res, err := commandRunner.Run("exit", "2")
	assert.NotNil(t, err)
	assert.Equal(t, "", string(res))
}
