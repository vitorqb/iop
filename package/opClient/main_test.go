package opClient

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/accountStorage"
	"github.com/vitorqb/iop/package/opClient/commandRunner"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tempFiles"
	"github.com/vitorqb/iop/package/testUtils"
	"github.com/vitorqb/iop/package/tokenStorage"
)

func TestNewCreatesANewClientInstance(t *testing.T) {
	sys := system.NewMock()
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
	accountStorage := accountStorage.NewInMemoryAccountStorage("")
	commandRunner := commandRunner.NewMockedCommandRunner("", nil)
	opClient := New(&sys, &tokenStorage, &accountStorage, &commandRunner)
	if opClient.path != DEFAULT_CLIENT {
		t.Fatal("Unexpected path")
	}
}

func TestRunWithTokenAppendsToken(t *testing.T) {
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("FOO")
	commandRunner := commandRunner.CommandRunner{}
	opClient := NewTestOpClient(
		WithTokenStorage(&tokenStorage),
		WithPath("echo"),
		WithCommandRunner(commandRunner),
	)
	result, err := opClient.runWithToken("bar")
	assert.Nil(t, err)
	assert.Equal(t, string(result), "--session FOO bar\n")
}

func TestEnsureLoggedInReturnsTrueIfAlreadyLoggedIn(t *testing.T) {
	// The mocked cmd runner returns `nil` as error, which is simulates being
	// logged in for `whoami` command.
	commandRunner := commandRunner.NewMockedCommandRunner("", nil)
	opClient := NewTestOpClient(WithCommandRunner(&commandRunner))
	opClient.EnsureLoggedIn()
	assert.Equal(t, commandRunner.CallCount, 1)
}

func TestEnsureLoggedInCallsLoginIfNotLoggedIn(t *testing.T) {
	// The mocked cmd runner returns exit 1 as error, which is
	// simulates NOT being logged in for `whoami` command.
	exitOneErr := exec.Command("exit", "1").Run()
	commandRunner := commandRunner.NewMockedCommandRunner("", exitOneErr)
	opClient := NewTestOpClient(WithCommandRunner(&commandRunner))
	opClient.EnsureLoggedIn()
	assert.Equal(t, commandRunner.CallCount, 2)
}

func TestEnsureLoggedInSavesTokenUsingTokenStorage(t *testing.T) {
	templateData := struct{
		WhoAmIExitCode string
		Body string
	}{"1", "echo -n 123"}
	scriptPath := testUtils.RenderTemplateTestFile(t, "mocked_whoami_script.sh", templateData)
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("")
	opClient := NewTestOpClient(
		WithTokenStorage(&tokenStorage),
		WithPath(scriptPath),
		WithCommandRunner(commandRunner.CommandRunner{}),
	)
	opClient.EnsureLoggedIn()
	token, _ := opClient.getToken()
	assert.Equal(t, token, "123")
	assert.Equal(t, tokenStorage.Token, "123")
}

func TestEnsureLoggedInExitsIfCmdFails(t *testing.T) {
	mockSystem := system.NewMock(system.WMockPin("PIN"))
	err := tempFiles.NewTempScript("#!/bin/bash \nexit 1").Run(func(scriptPath string) {
		opClient := NewTestOpClient(
			WithSystem(&mockSystem),
			WithPath(scriptPath),
			WithCommandRunner(commandRunner.CommandRunner{}),
		)
		opClient.EnsureLoggedIn()
		assert.Equal(t, mockSystem.CrashCallCount, 1)
		assert.Equal(t, mockSystem.LastCrashErrMsg, "Something wen't wrong during signin")
	})
	if err != nil {
		t.Error(err)
	}
}

func TestEnsureLoggedInRunsCorrectCommand(t *testing.T) {
	// The mocked cmd runner returns exit 1 as error, which is
	// simulates NOT being logged in for `whoami` command.
	exitOneErr := exec.Command("exit", "1").Run()
	tokenStorage := tokenStorage.NewInMemoryTokenStorage("token")
	accountStorage := accountStorage.NewInMemoryAccountStorage("account")
	commandRunner := commandRunner.NewMockedCommandRunner("", exitOneErr)
	opClient := NewTestOpClient(
		WithTokenStorage(&tokenStorage),
		WithAccountStorage(&accountStorage),
		WithPath("echo"),
		WithCommandRunner(&commandRunner),
	)
	opClient.EnsureLoggedIn()
	assert.Equal(t, commandRunner.LastArgs, []string{"echo", "signin", "--raw", "--session", "token", "--account", "account"})
}

func TestGetPasswordRetunsThePassword(t *testing.T) {
	testDataFilePath, err := testUtils.GetTestDataFilePath("password_field_1.json")
	assert.Nil(t, err)
	testFileCatScript := tempFiles.NewTempCat(testDataFilePath)
	err = testFileCatScript.Run(func(scriptPath string) {
		opClient := NewTestOpClient(
			WithPath(scriptPath),
			WithCommandRunner(commandRunner.CommandRunner{}),
		)
		assert.Equal(t, opClient.GetPassword("itemRef"), "the-password")
	})
	assert.Nil(t, err)
}

func TestListItemTitlesReturnItemTitles(t *testing.T) {
	testDataFilePath, _ := testUtils.GetTestDataFilePath("op_list_1.json")
	expectedTitles := []string{"some title 1", "some title 2"}
	testFileCatScript := tempFiles.NewTempCat(testDataFilePath)
	err := testFileCatScript.Run(func(scriptPath string) {
		opClient := NewTestOpClient(
			WithPath(scriptPath),
			WithCommandRunner(commandRunner.CommandRunner{}),
		)
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
		opClient := NewTestOpClient(WithPath(scriptPath))
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
		err := tempFiles.NewTempScript(script).Run(func(path string) {
			opClient := NewTestOpClient(
				WithPath(path),
				WithCommandRunner(commandRunner.CommandRunner{}),
			)

			// ACT
			result, err := opClient.isLoggedIn()

			// ASSERT
			assert.Nil(t, err)
			assert.Equal(t, testCase.ExpectedResult, result, "%+v", testCase)
		})
		assert.Nil(t, err)		
	}

}
