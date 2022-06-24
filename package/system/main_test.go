package system

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/tempFiles"
	"github.com/vitorqb/iop/package/testUtils"
)

func TestAskUserToSelectStringReturnsSelected(t *testing.T) {
	err := tempFiles.NewTempScript("#!/bin/sh \nhead -n1").Run(func(scriptPath string) {
		system := System{userSelectProgram: []string{scriptPath}}
		result, err := system.AskUserToSelectString([]string{"atitle"})
		assert.Nil(t, err)
		assert.Equal(t, "atitle", result)
	})
	if err != nil {
		t.Error(err)
	}
}

func TestAskForUserPinHappyPath(t *testing.T) {
	fakePinEntryData := struct{Pin string}{"FOO"}
	fakePinEntry := testUtils.RenderTemplateTestFile(t, "fakePinEntry.sh", fakePinEntryData)
	system := System{pinentryProgram: []string{fakePinEntry}}
	output, err := system.AskUserForPin("")
	assert.Nil(t, err)
	assert.Equal(t, "FOO", output)
}

func TestNotifyUserHappyPath(t *testing.T) {
	// TODO - Extract tempfile to tempFiles and remove old implementation
	tempFile, err := ioutil.TempFile("", "*")
	t.Cleanup(func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	})
	assert.Nil(t, err)
	fakeNotifySend := testUtils.RenderTemplateTestFile(
		t,
		"fakeNotifySend.sh",
		struct{OutputFile string}{tempFile.Name()},
	)
	system := System{notifySendProgram: fakeNotifySend}
	err = system.NotifyUser("A title", "A body")
	assert.Nil(t, err)

	content, err := ioutil.ReadFile(tempFile.Name())
	assert.Nil(t, err)
	assert.Equal(t, "title=A title;body=A body", string(content))
}
