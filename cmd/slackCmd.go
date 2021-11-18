package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var slackCmd = &cobra.Command{
	Use:   "slack",
	Short: "Notifications for Slack",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Not implemented yet...")
	},
}
