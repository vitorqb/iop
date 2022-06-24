package provider

import (
	"os"
	"path/filepath"

	"github.com/vitorqb/iop/package/accountStorage"
	"github.com/vitorqb/iop/package/opClient"
	"github.com/vitorqb/iop/package/opClient/commandRunner"
	"github.com/vitorqb/iop/package/storage"
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

func TokenStorage(system system.ISystem) storage.ISimpleStorage {
	tokenStorage, err := tokenStorage.New("")
	if err != nil {
		system.Crash("Could not initialize opClient", err)
	}
	return tokenStorage
}

func AccountStorage(system system.ISystem) accountStorage.IAccountStorage {
	homeDir := getUserDir(system)
	filePath := filepath.Join(homeDir, ".iop/currentAccount")
	return accountStorage.New(filePath)
}

func OpClient(
	sys system.ISystem,
	tokenStorage   storage.ISimpleStorage,
	accountStorage accountStorage.IAccountStorage,
) opClient.IOpClient {
	commandRunner := commandRunner.CommandRunner{}
	return opClient.New(sys, tokenStorage, accountStorage, commandRunner)
}
