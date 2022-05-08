package cmd

import (
	"log"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/vitorqb/iop/package/opClient"
	"github.com/vitorqb/iop/package/system"
	"github.com/vitorqb/iop/package/tokenStorage"
)

var copyPasswordCmd = &cobra.Command{
	Use:   "copy-password",
	Short: "Copies a password to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		system := system.New()
		tokenStorage, err := tokenStorage.New("")
		if err != nil {
			system.Crash("Could not initialize opClient", err)
		}
		itemRef, _ := cmd.Flags().GetString("ref")
		client := opClient.New(&system, tokenStorage)
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
		clipboard.WriteAll(password)
		log.Println("Copied to clipboard!")
	},
}

func init() {
	copyPasswordCmd.Flags().StringP("ref", "n", "", "Reference (name,id,link) of the item to copy.")

	rootCmd.AddCommand(copyPasswordCmd)
}
