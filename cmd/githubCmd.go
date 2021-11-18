package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var githubCmd = &cobra.Command{
	Use:   "github",
	Short: "Notifications for Github",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet...")
	},
}
