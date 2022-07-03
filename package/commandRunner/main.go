package commandRunner

import (
	"bytes"
	"io"
	"os/exec"
)

type ICommandRunner interface {
	Run(arg0 string, args ...string) ([]byte, error)
	RunWithStdin(stdin string, arg0 string, args ...string) ([]byte, error) // Runs passing a string as stdin
}

type CommandRunner struct {}
func (c CommandRunner) Run(arg0 string, args ...string) ([]byte, error) {
	bytes, err := exec.Command(arg0, args...).Output()
	return bytes, err
}
func (c CommandRunner) RunWithStdin(stdin string, arg0 string, args ...string) ([]byte, error) {
	cmd := exec.Command(arg0, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	err = cmd.Start()
	if err != nil {
		return nil, err
	}
	_, err = io.WriteString(stdinPipe, stdin + "\n")
	if err != nil {
		return nil, err
	}
	err = cmd.Wait()
	if err != nil {
		return nil, err
	}
	return stdout.Bytes(), err
}

// Mocked implementation for tests
type MockedCommandRunner struct {
	CallCount   int
	LastArgs    []string
	ReturnValue string
	Error       error
}
func NewMockedCommandRunner(returnValue string, error error) MockedCommandRunner {
	return MockedCommandRunner {
		CallCount: 0,
		LastArgs: []string{},
		Error: error,
		ReturnValue: returnValue,
	}
}
func (c *MockedCommandRunner) Run(arg0 string, args ...string) ([]byte, error) {
	c.CallCount++
	c.LastArgs = append([]string{arg0}, args...)
	return []byte(c.ReturnValue), c.Error
}
func (c *MockedCommandRunner) RunWithStdin(stdin string, arg0 string, args ...string) ([]byte, error) {
	return c.Run(arg0, args...)
}
