package bitbucket

import (
	"encoding/json"
	"fmt"
	"gonotify/internal/client"
)

const (
	bitbucketBuildStatusURI   string = "rest/build-status/1.0/commits"
	bitbucketBuildStatsURI    string = "rest/build-status/1.0/commits/stats"
	bitbucketCommitCommentURI string = "rest/api/1.0/projects/{projectKey}/repos/{repositorySlug}/commits/{commitId}/comments"
)

type BitbucketBuildState string

type BitbucketBuild struct {
	CommitId string
	Status   BitBucketBuildStatus
}

// BitBucketBuildStatus is used to set the payload
type BitBucketBuildStatus struct {
	Key         string `json:"key"`
	State       string `json:"state"`
	Url         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BitbucketClient struct {
	Host     string
	Username string
	Password string
	Build    BitbucketBuild
	Payload  interface{}
}

type CommitComment struct {
	Text string `json:"text"`
}

func (b BitbucketClient) BuildStatus() error {
	url := fmt.Sprintf("%s/%s/%s", b.Host, bitbucketBuildStatusURI, b.Build.CommitId)

	payload, err := json.Marshal(b.Build.Status)
	if err != nil {
		return err
	}

	creds := make(map[string]string)
	creds["username"] = b.Username
	creds["password"] = b.Password

	client.Post(url, payload, creds)

	comment := fmt.Sprintf("GONOTIFY: the build was marked as %s\n"+
		"Message:\n"+
		"```text\n"+
		"Replace this message!\n"+
		"```\n"+
		"JSON Response:\n"+
		"```json\n"+
		"{'msg':'Example Error','verbose':'this is an example error!'}\n"+
		"```", b.Build.Status.State)
	b.Comment(comment)
	return nil
}

func (b BitbucketClient) Comment(msg string) {
	url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/pull-requests/%s/comments", b.Host, b.Build.Status.Key, "root", "15")
	//url := fmt.Sprintf("%s/rest/api/1.0/projects/%s/repos/%s/commits/%s/comments", b.Host, b.Build.Status.Key, "root", b.Build.CommitId)
	message := CommitComment{
		Text: msg,
	}
	payload, _ := json.Marshal(message)
	creds := make(map[string]string)
	creds["username"] = b.Username
	creds["password"] = b.Password
	client.Post(url, payload, creds)
}

func (b BitBucketBuildStatus) Payload() ([]byte, error) {
	return json.Marshal(b)
}

func NewBitbucketBuild() BitbucketBuild {
	return BitbucketBuild{}
}

type JsonPayload interface {
	Payload() ([]byte, error)
}
