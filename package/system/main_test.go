package system

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vitorqb/iop/package/tempFiles"
)

func TestAskUserToSelectStringReturnsSelected(t *testing.T) {
	err := tempFiles.NewTempScript("#!/bin/sh \nhead -n1").Run(func(scriptPath string) {
		system := System{userSelectProgram: scriptPath}
		result, err := system.AskUserToSelectString([]string{"atitle"})
		assert.Nil(t, err)
		assert.Equal(t, "atitle", result)
	})
	if err != nil {
		t.Error(err)
	}
}
