package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gitlabCmd = &cobra.Command{
	Use:   "gitlab",
	Short: "Notifications for Gitlab",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet...")
	},
}
