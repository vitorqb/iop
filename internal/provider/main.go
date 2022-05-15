package provider

import (
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tokenStorage"
	"github.com/vitorqb/iop/package/opClient"
)

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

func OpClient(sys system.ISystem, tokenStorage tokenStorage.ITokenStorage) *opClient.OpClient {
	return opClient.New(sys, tokenStorage)
}
