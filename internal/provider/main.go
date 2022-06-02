package provider

import (
	"os"
	"path/filepath"

	"github.com/vitorqb/iop/package/emailStorage"
	"github.com/vitorqb/iop/package/opClient"
	"github.com/vitorqb/iop/package/opClient/commandRunner"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tokenStorage"
)

func getUserDir(system system.ISystem) string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		system.Crash("Could not get user home directory", err)
	}
	return homeDir
}

func System() system.ISystem {
	system := system.New()
	return &system
}

func TokenStorage(system system.ISystem) tokenStorage.ITokenStorage {
	tokenStorage, err := tokenStorage.New("")
	if err != nil {
		system.Crash("Could not initialize opClient", err)
	}
	return tokenStorage
}

func EmailStorage(system system.ISystem) emailStorage.IEmailStorage {
	homeDir := getUserDir(system)
	filePath := filepath.Join(homeDir, ".iop/currentAccount")
	return emailStorage.New(filePath)
}

func OpClient(
	sys system.ISystem,
	tokenStorage tokenStorage.ITokenStorage,
	emailStorage emailStorage.IEmailStorage,
) opClient.IOpClient {
	commandRunner := commandRunner.CommandRunner{}
	return opClient.New(sys, tokenStorage, emailStorage, commandRunner)
}
