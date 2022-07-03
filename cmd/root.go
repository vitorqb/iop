package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/vitorqb/pmwrap/cmd/copyPassword"
	"github.com/vitorqb/pmwrap/cmd/selectAccount"
	"github.com/vitorqb/pmwrap/internal/config"
)

var rootCmd = &cobra.Command{
	Use:   "pmwrap",
	Short: "Quickly copy passwords from password managers to the clipboard",
	Long:  `pmwrap fits the needs of users of password managers that really only care about querying for their passwords quickly and easily.`,
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
