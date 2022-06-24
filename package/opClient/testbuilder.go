package opClient

import (
	"github.com/vitorqb/iop/package/accountStorage"
	"github.com/vitorqb/iop/package/opClient/commandRunner"
	"github.com/vitorqb/iop/package/storage"
	"github.com/vitorqb/iop/package/system"
)

// Helper builder for OpClient
type opTestClientBUildOptions struct {
	tokenStorage    storage.ISimpleStorage
	sys             system.ISystem
	path            string
	accountStorage  storage.ISimpleStorage
	commandRunner   commandRunner.ICommandRunner
}
type opTestClientBuilder func(o *opTestClientBUildOptions)
func WithTokenStorage(t storage.ISimpleStorage) opTestClientBuilder {
	return func(o *opTestClientBUildOptions) {
		o.tokenStorage = t
	}
}
func WithPath(p string) opTestClientBuilder {
	return func(o *opTestClientBUildOptions) {
		o.path = p
	}
}
func WithCommandRunner(c commandRunner.ICommandRunner) opTestClientBuilder {
	return func(o *opTestClientBUildOptions) {
		o.commandRunner = c
	}
}
func WithAccountStorage(s storage.ISimpleStorage) opTestClientBuilder {
	return func(o *opTestClientBUildOptions) {
		o.accountStorage = s
	}
}
func WithSystem(s system.ISystem) opTestClientBuilder {
	return func(o *opTestClientBUildOptions) {
		o.sys = s
	}
}
func NewTestOpClient(builders ...opTestClientBuilder) OpClient {
	tokenStorage := storage.NewInMemoryTokenStorage("")
	mockSystem := system.NewMock()
	accountStorage := accountStorage.NewInMemoryAccountStorage("")
	commandRunner := commandRunner.MockedCommandRunner{ReturnValue: ""}
	options := opTestClientBUildOptions {
		tokenStorage:   &tokenStorage,
		sys:            &mockSystem,
		path:           "",
		accountStorage: &accountStorage,
		commandRunner:  &commandRunner,
	}
	for _, builder := range builders {
		builder(&options)
	}
	return OpClient(options)
}


