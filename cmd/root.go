package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vitorqb/iop/cmd/copyPassword"
	"github.com/vitorqb/iop/cmd/selectAccount"
	"github.com/vitorqb/iop/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "iop",
	Short: "An improved (or simplified) One Password client for mortals",
	Long:  `iop fits the needs of users of 1password that really only care about querying for their passwords quickly and easily.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string)  {
		err := config.LoadConfig()
		if err != nil {
			panic(fmt.Errorf("Error loading config: %w", err))
		}

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	err := config.LoadViper(rootCmd)
	if err != nil {
		panic(fmt.Errorf("Error setting up command: %w", err))
	}
	copyPassword.Setup(rootCmd)
	selectAccount.Setup(rootCmd)
}
