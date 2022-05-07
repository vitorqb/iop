package cmd

import (
	"log"

	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"

	"github.com/vitorqb/iop/package/opClient"
	"github.com/vitorqb/iop/package/system"
)

var copyPasswordCmd = &cobra.Command{
	Use:   "copy-password",
	Short: "Copies a password to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		itemRef, _ := cmd.Flags().GetString("ref")
		client := opClient.New()
		client.EnsureLoggedIn()

		// TODO -> If we init a system here, why not pass it to opClient?
		if itemRef == "" {
			system := system.New()
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
