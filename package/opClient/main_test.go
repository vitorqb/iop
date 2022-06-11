package opClient

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/accountStorage"
	"github.com/vitorqb/iop/package/opClient/commandRunner"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tempFiles"
	"github.com/vitorqb/iop/package/testUtils"
	"github.com/vitorqb/iop/package/tokenStorage"
)

// !!!! TODO Can we have a builder for opClient that defaults stuff?

func TestNewCreatesANewClientInstance(t *testing.T) {
	sys := system.NewMock()
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
	accountStorage := accountStorage.NewInMemoryAccountStorage("")
	commandRunner := commandRunner.MockedCommandRunner{ReturnValue: ""}
	opClient := New(&sys, &tokenStorage, &accountStorage, &commandRunner)
	if opClient.path != DEFAULT_CLIENT {
		t.Fatal("Unexpected path")
	}
}

func TestRunWithTokenAppendsToken(t *testing.T) {
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("FOO")
	accountStorage := accountStorage.NewInMemoryAccountStorage("")
	opClient := OpClient{
		tokenStorage: &tokenStorage,
		accountStorage: &accountStorage,
		path:         "echo",
		commandRunner: commandRunner.CommandRunner{},
	}
	result, err := opClient.runWithToken("bar")
	assert.Nil(t, err)
	assert.Equal(t, string(result), "--session FOO bar\n")
}

func TestEnsureLoggedInSavesTokenUsingTokenStorage(t *testing.T) {
	err := tempFiles.NewTempScript("#!/bin/sh \necho -n 123").Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		accountStorage := accountStorage.NewInMemoryAccountStorage("")
		opClient := OpClient{
			path:         scriptPath,
			tokenStorage: &tokenStorage,
			accountStorage: &accountStorage,
			commandRunner: commandRunner.CommandRunner{},
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
		accountStorage := accountStorage.NewInMemoryAccountStorage("")
		opClient := OpClient{
			tokenStorage: &tokenStorage,
			accountStorage: &accountStorage,
			sys:          &mockSystem,
			path:         scriptPath,
			commandRunner: commandRunner.CommandRunner{},
		}
		opClient.EnsureLoggedIn()
		assert.Equal(t, mockSystem.CrashCallCount, 1)
		assert.Equal(t, mockSystem.LastCrashErrMsg, "Something wen't wrong during signin")
	})
	if err != nil {
		t.Error(err)
	}
}

func TestEnsureLoggedInRunsCorrectCommand(t *testing.T) {
	mockSystem := system.NewMock()
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("token")
	accountStorage := accountStorage.NewInMemoryAccountStorage("account")
	commandRunner := commandRunner.MockedCommandRunner{ReturnValue: ""}
	opClient := OpClient{
		tokenStorage: &tokenStorage,
		accountStorage: &accountStorage,
		sys: &mockSystem,
		path: "echo",
		commandRunner: &commandRunner,
	}
	opClient.EnsureLoggedIn()
	assert.Equal(t, commandRunner.LastArgs, []string{"echo", "signin", "--raw", "--session", "token", "--account", "account"})
}

func TestGetPasswordRetunsThePassword(t *testing.T) {
	testDataFilePath, err := testUtils.GetTestDataFilePath("password_field_1.json")
	if err != nil {
		t.Error(err)
	}
	testFileCatScript := tempFiles.NewTempCat(testDataFilePath)
	err = testFileCatScript.Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		accountStorage := accountStorage.NewInMemoryAccountStorage("")
		opClient := OpClient{
			tokenStorage: &tokenStorage,
			accountStorage: &accountStorage,
			path:         scriptPath,
			commandRunner: commandRunner.CommandRunner{},
		}
		assert.Equal(t, opClient.GetPassword("itemRef"), "the-password")
	})
	if err != nil {
		t.Error(err)
	}
}

func TestListItemTitlesReturnItemTitles(t *testing.T) {
	testDataFilePath, _ := testUtils.GetTestDataFilePath("op_list_1.json")
	expectedTitles := []string{"some title 1", "some title 2"}
	testFileCatScript := tempFiles.NewTempCat(testDataFilePath)
	err := testFileCatScript.Run(func(scriptPath string) {
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
		accountStorage := accountStorage.NewInMemoryAccountStorage("")
		opClient := OpClient{
			tokenStorage: &tokenStorage,
			accountStorage: &accountStorage,
			path:         scriptPath,
			commandRunner: commandRunner.CommandRunner{},
		}
		assert.ElementsMatch(t, expectedTitles, opClient.ListItemTitles())
	})
	if err != nil {
		t.Error(err)
	}
}

func TestListAccountsReturnAccounts(t *testing.T) {
	testDataFilePath, _ := testUtils.GetTestDataFilePath("op_accounts_list_1.json")
	expectedAccounts := []string{"team_foo", "my"}
	testFileCatScript := tempFiles.NewTempCat(testDataFilePath)
	err := testFileCatScript.Run(func(scriptPath string) {
		opClient := OpClient{ path: scriptPath }
		accounts, err := opClient.ListAccounts()
		if err != nil {
			t.Error(err)
		}
		assert.ElementsMatch(t, expectedAccounts, accounts)
	})
	if err != nil {
		t.Error(err)
	}	
}

func TestIsLoggedInTrue(t *testing.T) {

	testCases := []struct{
		ExitCode string
		ExpectedResult bool
	}{
		{"0", true},
		{"1", false},
	}

	for _, testCase := range testCases {
		// ARRANGE
		script := fmt.Sprintf("#!/bin/bash\nexit %s", testCase.ExitCode)
		tokenStorage := tokenStorage.NewInMemoryTokenStorage("foo")
		accountStorage := accountStorage.NewInMemoryAccountStorage("bar")
		err := tempFiles.NewTempScript(script).Run(func(path string) {
			opClient := OpClient{
				tokenStorage:   &tokenStorage,
				accountStorage: &accountStorage,
				path:           path,
				commandRunner:  commandRunner.CommandRunner{},
			}

			// ACT
			result, err := opClient.isLoggedIn()

			// ASSERT
			assert.Nil(t, err)
			assert.Equal(t, testCase.ExpectedResult, result, "%+v", testCase)
		})
		assert.Nil(t, err)		
	}

}
