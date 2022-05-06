package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/atotto/clipboard"
	
	"github.com/vitorqb/iop/package/opClient"
)

var copyPasswordCmd = &cobra.Command{
	Use: "copy-password",
	Short: "Copies a password to clipboard",
	Run: func(cmd *cobra.Command, args []string) {
		itemRef, _ := cmd.Flags().GetString("ref")
		client := opClient.New()
		client.EnsureLoggedIn()
		password := client.GetPassword(itemRef)
		clipboard.WriteAll(password)
		log.Println("Copied to clipboard!")
	},
}

func init() {
	copyPasswordCmd.Flags().StringP("ref", "n", "", "Reference (name,id,link) of the item to copy.")
	copyPasswordCmd.MarkFlagRequired("ref")

	rootCmd.AddCommand(copyPasswordCmd)
}
