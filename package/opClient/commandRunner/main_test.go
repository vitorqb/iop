package commandRunner

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/pmwrap/package/testUtils"
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

func TestCommandRunnerRunWithStdin(t *testing.T) {
	testScript := testUtils.RenderTemplateTestFile(t, "testscript.sh", nil)
	output, err := CommandRunner{}.RunWithStdin("FOO", testScript, "--bar")
	assert.Nil(t, err)
	assert.Equal(t, "arg1=--bar,arg2=,stdin=FOO\n", string(output))
}
