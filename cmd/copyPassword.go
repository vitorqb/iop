package cmd

import (
	"log"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/vitorqb/iop/internal/provider"
)

var copyPasswordCmd = &cobra.Command{
	Use:   "copy-password",
	Short: "Copies a password to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		system := provider.System()
		tokenStorage := provider.TokenStorage(system)
		client := provider.OpClient(system, tokenStorage)

		itemRef, _ := cmd.Flags().GetString("ref")

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
		log.Println("Copied to clipboard!")
	},
}

func init() {
	copyPasswordCmd.Flags().StringP("ref", "n", "", "Reference (name,id,link) of the item to copy.")

	rootCmd.AddCommand(copyPasswordCmd)
}
