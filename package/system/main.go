package system

import (
	"log"
	"os"
)

// An interface to control the System, mostly useful for tests
type ISystem interface {
	Crash(errMsg string, err error)
}

// A real system implementation
type System struct{}

func (s *System) Crash(errMsg string, err error) {
	log.Fatal(errMsg, " - ", err)
	os.Exit(99)
}
func New() System {
	return System{}
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
func NewMock() MockSystem {
	return MockSystem{CrashCallCount: 0}
}
