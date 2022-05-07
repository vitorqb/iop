package system

import (
	"log"
	"os"
)

// An interface to control the System, mostly useful for tests
type ISystem interface {
	Crash(errMsg string, err error)
	exit(code int)
}

// A real system implementation
type System struct{}

func (s *System) Crash(errMsg string, err error) {
	log.Fatal(errMsg, " - ", err)
	s.exit(99)
}
func (s *System) exit(code int) {
	os.Exit(code)
}
func New() System {
	return System{}
}

// A mock system for test
type MockSystem struct {
	CrashCallCount int
	LastCrashErr error
	LastCrashErrMsg string
}

func (s *MockSystem) Crash(errMsg string, err error) {
	s.CrashCallCount++
	s.LastCrashErr = err
	s.LastCrashErrMsg = errMsg
}
func (s *MockSystem) exit(code int) {}
func NewMock() MockSystem {
	return MockSystem{CrashCallCount: 0}
}
