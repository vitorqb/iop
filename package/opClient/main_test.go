package opClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/emailStorage"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tempFiles"
	"github.com/vitorqb/iop/package/testUtils"
	"github.com/vitorqb/iop/package/tokenStorage"
)

// !!!! TODO Can we have a builder for opClient that defaults stuff?

func TestNewCreatesANewClientInstance(t *testing.T) {
	sys := system.NewMock()
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
	emailStorage := emailStorage.NewInMemoryEmailStorage("")
	opClient := New(&sys, &tokenStorage, &emailStorage)
	if opClient.path != DEFAULT_CLIENT {
		t.Fatal("Unexpected path")
	}
}

func TestRunWithTokenAppendsToken(t *testing.T) {
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("FOO")
	emailStorage := emailStorage.NewInMemoryEmailStorage("")
	opClient := OpClient{
		tokenStorage: &tokenStorage,
		emailStorage: &emailStorage,
		path:         "echo",
	}
	result, err := opClient.runWithToken("bar")
	assert.Nil(t, err)
	assert.Equal(t, string(result), "--session FOO bar\n")
}

func TestEnsureLoggedInSavesTokenUsingTokenStorage(t *testing.T) {
	err := tempFiles.NewTempScript("#!/bin/sh \necho -n 123").Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		emailStorage := emailStorage.NewInMemoryEmailStorage("")
		opClient := OpClient{
			path:         scriptPath,
			tokenStorage: &tokenStorage,
		emailStorage: &emailStorage,
		}
		opClient.EnsureLoggedIn()
		token, _ := opClient.getToken()
		assert.Equal(t, token, "123")
		assert.Equal(t, tokenStorage.Token, "123")
	})
	if err != nil {
		t.Error(err)
	}
}

func TestEnsureLoggedInExitsIfCmdFails(t *testing.T) {
	mockSystem := system.NewMock()
	err := tempFiles.NewTempScript("#!/bin/bash \nexit 1").Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		emailStorage := emailStorage.NewInMemoryEmailStorage("")
		opClient := OpClient{
			tokenStorage: &tokenStorage,
		emailStorage: &emailStorage,
			sys:          &mockSystem,
			path:         scriptPath,
		}
		opClient.EnsureLoggedIn()
		assert.Equal(t, mockSystem.CrashCallCount, 1)
		assert.Equal(t, mockSystem.LastCrashErrMsg, "Something wen't wrong during signin")
	})
	if err != nil {
		t.Error(err)
	}
}

func TestGetPasswordRetunsThePassword(t *testing.T) {
	err := tempFiles.NewTempScript("#!/bin/sh \necho -n '12345\n'").Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		emailStorage := emailStorage.NewInMemoryEmailStorage("")
		opClient := OpClient{
			tokenStorage: &tokenStorage,
		emailStorage: &emailStorage,
			path:         scriptPath,
		}
		assert.Equal(t, opClient.GetPassword("itemRef"), "12345")
	})
	if err != nil {
		t.Error(err)
	}
}

func TestListItemTitlesReturnItemTitles(t *testing.T) {
	testDataFilePath, _ := testUtils.GetTestDataFilePath("op_list_1.json")
	expectedTitles := []string{"some title 1", "some title 2"}
	testFileCatScript := tempFiles.NewTempScript("#!/bin/sh \ncat " + testDataFilePath)
	err := testFileCatScript.Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		emailStorage := emailStorage.NewInMemoryEmailStorage("")
		opClient := OpClient{
			tokenStorage: &tokenStorage,
		emailStorage: &emailStorage,
			path:         scriptPath,
		}
		assert.ElementsMatch(t, expectedTitles, opClient.ListItemTitles())
	})
	if err != nil {
		t.Error(err)
	}
}

func TestListEmailsReturnEmails(t *testing.T) {
	testDataFilePath, _ := testUtils.GetTestDataFilePath("op_accounts_list_1.json")
	expectedEmails := []string{"antonio.bababa@foo.com", "antonioqb@gmail.com"}
	testFileCatScript := tempFiles.NewTempScript("#!/bin/sh \ncat " + testDataFilePath)
	err := testFileCatScript.Run(func(scriptPath string) {
		opClient := OpClient{ path: scriptPath }
		emails, err := opClient.ListEmails()
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, expectedEmails, emails)
	})
	if err != nil {
		t.Error(err)
	}	
}
