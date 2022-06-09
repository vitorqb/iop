package system

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/vitorqb/iop/internal/config"
)

// An interface to control the System
type ISystem interface {
	Crash(errMsg string, err error)
	AskUserToSelectString(options []string) (string, error)
}

// A real system implementation
type System struct {
	userSelectProgram []string
}

func (s *System) Crash(errMsg string, err error) {
	log.Fatal(errMsg, " - ", err)
	os.Exit(99)
}
func (s *System) AskUserToSelectString(options []string) (string, error) {
	stdinBuff := bytes.Buffer{}
	for _, string := range options {
		stdinBuff.WriteString(string + "\n")
	}
	dmenu := exec.Command(s.userSelectProgram[0], s.userSelectProgram[1:]...)
	dmenu.Stdin = &stdinBuff
	resultInBytes, err := dmenu.Output()
	return strings.Trim(string(resultInBytes), "\n"), err
}
func New() System {
	config := config.GetConfig()
	return System{
		userSelectProgram: config.DmenuCommand,
	}
}

// A mock system for test
type MockSystem struct {
	CrashCallCount  int
	LastCrashErr    error
	LastCrashErrMsg string
}

func (s *MockSystem) Crash(errMsg string, err error) {
	s.CrashCallCount++
	s.LastCrashErr = err
	s.LastCrashErrMsg = errMsg
}
func (s *MockSystem) AskUserToSelectString(options []string) (string, error) {
	return options[0], nil
}
func NewMock() MockSystem {
	return MockSystem{CrashCallCount: 0}
}
