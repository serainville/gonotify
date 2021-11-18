package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	projectKey       string
	commitId         string
	buildState       string
	buildUrl         string
	buildName        string
	buildDescription string
	username         string
	password         string
	baseURL          string
)

var rootCmd = &cobra.Command{
	Use:   "gonotify",
	Short: "Send notifications!",
	Long:  "Gonotify is a tool that simplifies sending notifications to popular CI tools via their REST APIs.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(bitbucketCmd)
	rootCmd.AddCommand(githubCmd)
	rootCmd.AddCommand(gitlabCmd)
	rootCmd.AddCommand(slackCmd)
	rootCmd.AddCommand(versionCmd)

	bitbucketCmd.AddCommand(bitbucketBuildCmd)
	bitbucketCmd.AddCommand(bitbucketStatsCmd)

	bitbucketBuildCmd.Flags().StringVarP(&projectKey, "projectkey", "k", "", "Bitbucket project key")
	bitbucketBuildCmd.Flags().StringVarP(&commitId, "commitid", "c", "", "Commit ID")
	bitbucketBuildCmd.Flags().StringVarP(&buildUrl, "url", "l", "", "Build URL")
	bitbucketBuildCmd.Flags().StringVarP(&buildState, "state", "s", "INPROGRESS", "Build state (INPROGRESS|SUCCESSFUL|FAILED")
	bitbucketBuildCmd.Flags().StringVarP(&buildName, "name", "n", "", "Build status name")
	bitbucketBuildCmd.Flags().StringVarP(&buildDescription, "description", "d", "", "Description of build status")
	bitbucketCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "Basic Auth Username")
	bitbucketCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "Basic Auth Password")
	bitbucketCmd.PersistentFlags().StringVarP(&baseURL, "baseurl", "b", "", "Base URL of Bitbucket Server")

	//bitbucketBuildCmd.MarkFlagRequired("commitid")
	bitbucketBuildCmd.MarkFlagRequired("projectkey")
	bitbucketBuildCmd.MarkFlagRequired("url")
	bitbucketBuildCmd.MarkFlagRequired("state")
}
