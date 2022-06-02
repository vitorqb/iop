package commandRunner

import (
	"bytes"
	"os"
	"os/exec"
)

type ICommandRunner interface {
	Run(arg0 string, args ...string) ([]byte, error)
	RunAsProxy(arg0 string, args ...string)        ([]byte, error) // Runs leaving stdin and stdout from parent
}

type CommandRunner struct {}
func (c CommandRunner) Run(arg0 string, args ...string) ([]byte, error) {
	bytes, err := exec.Command(arg0, args...).Output()
	return bytes, err
}
func (c CommandRunner) RunAsProxy(arg0 string, args ...string) ([]byte, error) {
	cmd := exec.Command(arg0, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	return stdout.Bytes(), err
}

// Mocked implementation for tests
type MockedCommandRunner struct {
	LastArgs    []string
	ReturnValue string
}
func (c *MockedCommandRunner) Run(arg0 string, args ...string) ([]byte, error) {
	c.LastArgs = append([]string{arg0}, args...)
	return []byte(c.ReturnValue), nil
}
func (c *MockedCommandRunner) RunAsProxy(arg0 string, args ...string) ([]byte, error) {
	c.LastArgs = append([]string{arg0}, args...)
	return []byte(c.ReturnValue), nil
}
