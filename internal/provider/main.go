package provider

import (
	"os"
	"path/filepath"

	"github.com/vitorqb/pmwrap/package/accountStorage"
	"github.com/vitorqb/pmwrap/package/opClient"
	"github.com/vitorqb/pmwrap/package/opClient/commandRunner"
	"github.com/vitorqb/pmwrap/package/system"
	"github.com/vitorqb/pmwrap/package/tokenStorage"
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

func AccountStorage(system system.ISystem) accountStorage.IAccountStorage {
	homeDir := getUserDir(system)
	filePath := filepath.Join(homeDir, ".pmwrap/currentAccount")
	return accountStorage.New(filePath)
}

func OpClient(
	sys system.ISystem,
	tokenStorage tokenStorage.ITokenStorage,
	accountStorage accountStorage.IAccountStorage,
) opClient.IOpClient {
	commandRunner := commandRunner.CommandRunner{}
	return opClient.New(sys, tokenStorage, accountStorage, commandRunner)
}
