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

	var mappedEvents []events.MappedEvent

	for _, event := range theList {
		mappedEvents = append(mappedEvents, events.MappedEvent{
			Type:     event.Type,
			RepoName: event.Repo.Name,
		})
	}

	counts := make(map[events.MappedEvent]int, len(mappedEvents))

	for _, e := range mappedEvents {
		key := events.MappedEvent{
			Type:     e.Type,
			RepoName: e.RepoName,
		}

		counts[key]++
	}

	for k, v := range counts {
		switch k.Type {
		case "PushEvent":
			fmt.Printf("--- Pushed commits to a %v: %v times\n", k.RepoName, v)
		case "CreateEvent":
			fmt.Printf("--- Created a new repo: %v\n", k.RepoName)
		case "IssuesEvent":
			fmt.Printf("--- Opened, Closed or Edited and Issue in repo: %v %v times\n", k.RepoName, v)
		case "IssueCommentEvent":
			fmt.Printf("--- Commented on an issue in repo: %v  %v times\n", k.RepoName, v)
		case "PullRequestEvent":
			fmt.Printf("--- Opened, Closed or Merged a pull request in repo: %v %v times\n", k.RepoName, v)
		case "WatchEvent":
			fmt.Printf("--- Starred a repo: %v\n", k.RepoName)
		case "ReleaseEvent":
			fmt.Printf("--- Published a release in a repo: %v\n", k.RepoName)
		}
	}
}
