package selectAccount

import (
	"github.com/spf13/cobra"

	"github.com/vitorqb/pmwrap/internal/provider"
	"github.com/vitorqb/pmwrap/package/accountStorage"
	"github.com/vitorqb/pmwrap/package/opClient"
	"github.com/vitorqb/pmwrap/package/system"
)

var selectAccountCmd = &cobra.Command{
	Use:   "select-account",
	Short: "Selects the current active account",
	Run: func(cmd *cobra.Command, args []string) {
		system := provider.System()
		tokenStorage := provider.TokenStorage(system)
		accountStorage := provider.AccountStorage(system)
		client := provider.OpClient(system, tokenStorage, accountStorage)
		run(client, system, accountStorage)
	},
}

func run(
	client opClient.IOpClient,
	system system.ISystem,
	accountStorage accountStorage.IAccountStorage,
) {
	allAccounts, err := client.ListAccounts()
	if err != nil || len(allAccounts) == 0 {
		system.Crash("Failed to list accounts", err)
	}
	selectedAccount, err := system.AskUserToSelectString(allAccounts)
	if err != nil || selectedAccount == "" {
		system.Crash("Failed to ask user for account", err)
	}
	err = accountStorage.Put(selectedAccount)
	if err != nil {
		system.Crash("Failed to store current account", err)
	}
	_ = system.NotifyUser("PMWRAP", "Selected account: " + selectedAccount)
}

func Setup(rootCmd *cobra.Command) {
	// !!!! TODO ALLOW USER TO PASS FROM CLI ARGUMENT THE ACCOUNT
	rootCmd.AddCommand(selectAccountCmd)
}
