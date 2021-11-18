package cmd

import (
	"fmt"
	"gonotify/internal/bitbucket"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var bitbucketCmd = &cobra.Command{
	Use:   "bitbucket",
	Short: "Notifications for Atlassian Bitbucket Server",
}

var bitbucketBuildCmd = &cobra.Command{
	Use: "build",
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args[0]) == 0 || args[0] == "" {
			fmt.Println("A commit ID is required...")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := bitbucket.BitbucketClient{
			Host:     baseURL,
			Username: username,
			Password: password,
			Build: bitbucket.BitbucketBuild{
				CommitId: args[0],
				Status: bitbucket.BitBucketBuildStatus{
					Key:         projectKey,
					State:       buildState,
					Url:         buildUrl,
					Name:        buildName,
					Description: buildDescription,
				},
			},
		}
		client.BuildStatus()
	},
}

var bitbucketStatsCmd = &cobra.Command{
	Use: "stats",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("%s/%s/%s", "http://10.0.0.115:7990", "rest/build-status/1.0/commits/stats", "9129c4c5c92e85d7a971e413b157037109117ac7")

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Println("ERROR:", err)
			os.Exit(1)
		}

		req.SetBasicAuth("srainville", "NjEzOTk0NTgwNjg4Oi8U6k8sE05Y/bPIrWNMjiVDP4uM")

		fmt.Println(req.URL.String())

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)

	},
}
