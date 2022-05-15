package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "iop",
	Short: "An improved (or simplified) One Password client for mortals",
	Long:  `iop fits the needs of users of 1password that really only care about querying for their passwords quickly and easily.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
