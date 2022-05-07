package opClient

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tempScript"
	"github.com/vitorqb/iop/package/testUtils"
)

func TestNewCreatesANewClientInstance(t *testing.T) {
	opClient := New()
	if *opClient.token != "" {
		t.Fatal("Unexpected token")
	}
	if opClient.path != DEFAULT_CLIENT {
		t.Fatal("Unexpected path")
	}
}

func TestRunWithTokenAppendsToken(t *testing.T) {
	token := "token"
	opClient := OpClient{
		token: &token,
		path:  "echo",
	}
	result, err := opClient.runWithToken("bar")
	assert.Nil(t, err)
	assert.Equal(t, string(result), "--session "+token+" bar\n")
}

func TestEnsureLoggedInSetsToken(t *testing.T) {
	tempScript.New("#!/bin/sh \necho -n 123").Run(func(scriptPath string) {
		token := ""
		opClient := OpClient{
			token: &token,
			path:  scriptPath,
		}
		opClient.EnsureLoggedIn()
		assert.Equal(t, *opClient.token, "123")
	})
}

func TestEnsureLoggedInExitsIfCmdFails(t *testing.T) {
	mockSystem := system.NewMock()
	tempScript.New("#!/bin/bash \nexit 1").Run(func(scriptPath string) {
		token := ""
		opClient := OpClient{
			sys:   &mockSystem,
			token: &token,
			path:  scriptPath,
		}
		opClient.EnsureLoggedIn()
		assert.Equal(t, mockSystem.CrashCallCount, 1)
		assert.Equal(t, mockSystem.LastCrashErrMsg, "Something wen't wrong during signin")
	})
}

func TestGetPasswordRetunsThePassword(t *testing.T) {
	tempScript.New("#!/bin/sh \necho -n '12345\n'").Run(func(scriptPath string) {
		token := ""
		opClient := OpClient{
			token: &token,
			path:  scriptPath,
		}
		assert.Equal(t, opClient.GetPassword("itemRef"), "12345")
	})
}

func TestListItemTitlesReturnItemTitles(t *testing.T) {
	testDataFilePath, _ := testUtils.GetTestDataFilePath("op_list_1.json")
	expectedTitles := []string{"some title 1", "some title 2"}
	testFileCatScript := tempScript.New("#!/bin/sh \ncat " + testDataFilePath)
	testFileCatScript.Run(func(scriptPath string) {
		token := ""
		opClient := OpClient{
			token: &token,
			path:  scriptPath,
		}
		assert.ElementsMatch(t, expectedTitles, opClient.ListItemTitles())
	})
}
