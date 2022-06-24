package provider

import (
	"os"
	"path/filepath"

	"github.com/vitorqb/iop/internal/config"
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
	config := config.GetConfig()
	system := system.New(
		config.DmenuCommand,
		config.PinEntryCommand,
		config.NotifySendCommand,
	)
	return &system
}

func TokenStorage(system system.ISystem) storage.ISimpleStorage {
	tokenStorage, err := tokenStorage.New("")
	if err != nil {
		system.Crash("Could not initialize opClient", err)
	}
	return tokenStorage
}

func AccountStorage(system system.ISystem) storage.ISimpleStorage {
	homeDir := getUserDir(system)
	filePath := filepath.Join(homeDir, ".iop/currentAccount")
	return accountStorage.New(filePath)
}

func OpClient(
	sys system.ISystem,
	tokenStorage   storage.ISimpleStorage,
	accountStorage storage.ISimpleStorage,
) opClient.IOpClient {
	commandRunner := commandRunner.CommandRunner{}
	return opClient.New(sys, tokenStorage, accountStorage, commandRunner)
}
