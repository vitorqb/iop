package copyPassword

import (
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/vitorqb/iop/internal/provider"
	"github.com/vitorqb/iop/package/opClient"
	"github.com/vitorqb/iop/package/system"
)

var copyPasswordCmd = &cobra.Command{
	Use:   "copy-password",
	Short: "Copies a password to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		system := provider.System()
		tokenStorage := provider.TokenStorage(system)
		accountStorage := provider.AccountStorage(system)
		client := provider.OpClient(system, tokenStorage, accountStorage)

		itemRef, _ := cmd.Flags().GetString("ref")
		run(client, itemRef, system)
	},
}

func run(client opClient.IOpClient, itemRef string, system system.ISystem) {
	client.EnsureLoggedIn()
	if itemRef == "" {
		titles := client.ListItemTitles()
		selectedTitle, err := system.AskUserToSelectString(titles)
		if err != nil {
			system.Crash("Something went wrong asking user to select title", err)
		}
		itemRef = selectedTitle
	}

	password := client.GetPassword(itemRef)
	err := clipboard.WriteAll(password)
	if err != nil {
		system.Crash("Something went wrong when writing to clipboard", err)
	}
	_ = system.NotifyUser("IOP", "The password has been saved to clipboard and is ready to be used :)")
}

func Setup(rootCmd *cobra.Command) {
	copyPasswordCmd.Flags().StringP("ref", "n", "", "Reference (name,id,link) of the item to copy.")
	rootCmd.AddCommand(copyPasswordCmd)
}
