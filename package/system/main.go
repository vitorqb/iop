package system

import (
	"bufio"
	"bytes"
	"errors"
	"io"
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
	AskUserForPin(prompt string) (string, error)
}

// A real system implementation
type System struct {
	userSelectProgram []string
	pinentryProgram []string
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
func (s *System) AskUserForPin(prompt string) (string, error) {
	pinentry := exec.Command(s.pinentryProgram[0], s.pinentryProgram[1:]...)
	stdoutReader, stdoutWriter := io.Pipe()
	pinentry.Stdout = stdoutWriter
	stdin, err := pinentry.StdinPipe()
	if err != nil {
		return "", err
	}
	err = pinentry.Start()
	if err != nil {
		return "", err
	}
	defer func() {
		err := stdoutWriter.Close()
		if err != nil {
			log.Fatal("Error closing stdoutWrier: %w", err)
		}
		err = stdoutReader.Close()
		if err != nil {
			log.Fatal("Error closing stdoutReader: %w", err)
		}
		err = pinentry.Process.Kill()
		if err != nil {
			log.Fatal("Error closing pinenry: %w", err)
		}

	}()
	_, err = io.WriteString(stdin, "SETPOMPT " + prompt + "\n")
	if err != nil {
		return "", err
	}
	_, err = io.WriteString(stdin, "GETPIN\n")
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(stdoutReader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "D") {
			return strings.Trim(line[2:], "\n"), nil
		}
	}
	return "", errors.New("Failed to query usr for pin!")
}
func New() System {
	config := config.GetConfig()
	return System{
		userSelectProgram: config.DmenuCommand,
		pinentryProgram: config.PinEntryCommand,
	}
}

// A mock system for test
type MockSystem struct {
	CrashCallCount  int
	LastCrashErr    error
	LastCrashErrMsg string
	Pin             string
}

func (s *MockSystem) Crash(errMsg string, err error) {
	s.CrashCallCount++
	s.LastCrashErr = err
	s.LastCrashErrMsg = errMsg
}
func (s *MockSystem) AskUserToSelectString(options []string) (string, error) {
	return options[0], nil
}
func (s *MockSystem) AskUserForPin(prompt string) (string, error) {
	if (s.Pin == "") {
		return "", errors.New("[MockSystem] Failed to get pin.")
	}
	return s.Pin, nil
}
type MockOption func(s *MockSystem)
func WMockPin(pin string) MockOption {
	return func(s *MockSystem) {
		s.Pin = pin
	}
}
func NewMock(mockOptions ...MockOption) MockSystem {
	mockSystem := MockSystem{CrashCallCount: 0}
	for _, mockOption := range mockOptions {
		mockOption(&mockSystem)
	}
	return mockSystem
}
