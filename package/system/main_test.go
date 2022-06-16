package system

import (
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
