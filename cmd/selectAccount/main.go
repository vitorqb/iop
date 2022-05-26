package selectAccount

import (
	"github.com/spf13/cobra"

	"github.com/vitorqb/iop/internal/provider"
	"github.com/vitorqb/iop/package/emailStorage"
	"github.com/vitorqb/iop/package/opClient"
	"github.com/vitorqb/iop/package/system"
)

var selectAccountCmd = &cobra.Command{
	Use:   "select-account",
	Short: "Selects the current active account",
	Run: func(cmd *cobra.Command, args []string) {
		system := provider.System()
		tokenStorage := provider.TokenStorage(system)
		emailStorage := provider.EmailStorage(system)
		client := provider.OpClient(system, tokenStorage)
		run(client, system, emailStorage)
	},
}

func run(
	client opClient.IOpClient,
	system system.ISystem,
	emailStorage emailStorage.IEmailStorage,
) {
	allEmails, err := client.ListEmails()
	if err != nil || len(allEmails) == 0 {
		system.Crash("Failed to list accounts emails", err)
	}
	selectedEmail, err := system.AskUserToSelectString(allEmails)
	if err != nil || selectedEmail == "" {
		system.Crash("Failed to ask user for email", err)
	}
	err = emailStorage.Put(selectedEmail)
	if err != nil {
		system.Crash("Failed to store current email", err)
	}
	// !!!! TODO LOG SMTH TO THE USER (add to ISystem)
}

func Setup(rootCmd *cobra.Command) {
	// !!!! TODO ALLOW USER TO PASS FROM CLI ARGUMENT THE ACCOUNT
	rootCmd.AddCommand(selectAccountCmd)
}
