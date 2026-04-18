package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github-user-activity/events"
)

func main() {
	cliArguments := os.Args[1:]
	username := strings.Join(cliArguments, "")

	url := fmt.Sprintf("https://api.github.com/users/%v/events", username)

	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var theList []events.Event

	json.Unmarshal(body, &theList)

	// jsonFile, err := os.OpenFile("json.json", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// defer jsonFile.Close()

	// Happens when someone pushes commits to a repo
	pushCount := 0
	// Happens when someone creates a new repo
	createCount := 0
	// Triggered when an issue is: opened, closed, edited
	issueCount := 0
	// Someone commented on an issue
	commentedIssueCount := 0
	// A pull request was: opened, closed, merged
	pullCount := 0
	// Someone starred a repo
	watchCount := 0
	// A release was published in a repo
	releaseCount := 0

	for _, event := range theList {
		switch event.Type {
		case "PushEvent":
			pushCount++
		case "CreateEvent":
			createCount++
		case "IssuesEvent":
			issueCount++
		case "IssueCommentEvent":
			commentedIssueCount++
		case "PullRequestEvent":
			pullCount++
		case "WatchEvent":
			watchCount++
		case "ReleaseEvent":
			releaseCount++
		}
	}

	fmt.Printf("%v:\n", username)

	pushMsg := fmt.Sprintf("--- Pushed commits to a repo: %v times", pushCount)
	createMsg := fmt.Sprintf("--- Created a new repo: %v times", createCount)
	issueMsg := fmt.Sprintf("--- Opened, Closed or Edited and Issue: %v times", issueCount)
	commentIssueMsg := fmt.Sprintf("--- Commented on an issue: %v times", commentedIssueCount)
	pullMsg := fmt.Sprintf("--- Opened, Closed or Merged a pull request: %v times", pullCount)
	watchMsg := fmt.Sprintf("--- Starred a repo: %v times", watchCount)
	releaseMsg := fmt.Sprintf("--- Published a release in a repo: %v times", releaseCount)

	PrintMessage(pushCount, pushMsg)
	PrintMessage(createCount, createMsg)
	PrintMessage(issueCount, issueMsg)
	PrintMessage(commentedIssueCount, commentIssueMsg)
	PrintMessage(pullCount, pullMsg)
	PrintMessage(watchCount, watchMsg)
	PrintMessage(releaseCount, releaseMsg)
}

func PrintMessage(eventCount int, message string) {
	if eventCount > 0 {
		fmt.Println(message)
	}
}
