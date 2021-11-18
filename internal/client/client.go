package client

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Post(url string, payload []byte, creds map[string]string) {

	log.Println("Sending notification to Bitbucket Server...")
	log.Println("URL:", url)
	log.Println("Payload:", string(payload))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	req.SetBasicAuth(creds["username"], creds["password"])

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

	if OK := requestOK(resp.Status); !OK {
		log.Printf("Notification error encountered! (STATUS: %s)\n", strings.TrimSpace(resp.Status))
		return
	}
	log.Printf("Notification sent successfully (STATUS: %s)\n", strings.TrimSpace(resp.Status))

}

func Get() {

}

func requestOK(status string) bool {
	statusInt, err := strconv.ParseInt(strings.TrimSpace(status), 10, 16)
	if err != nil {
		log.Println(err)
	}

	if statusInt >= 200 && statusInt <= 300 {
		return true
	}
	return false
}
