package main

import "fmt"
import "github.com/google/go-github/github"
import "golang.org/x/oauth2"
import "context"
import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Token string
	User string
}

func main() {
	fmt.Printf("You only yolo once \n")

	repos := fetchRepos()
	printRepos(repos)
}

func fetchRepos() []*github.Repository {
	config := readToken()

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, config.User, nil)

	if err != nil {
		// wut
	}
	if repos != nil {
		// wut
	}

	return repos
}

func readToken() Configuration {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
	if len(configuration.Token) < 1 {
		fmt.Printf("Token not imported or empty.")
	}
	if len(configuration.User) < 1 {
		fmt.Printf("User not imported or empty.")
	}
	return configuration
}

func printRepos(repos []*github.Repository) {
	for _, v:= range repos {
		fmt.Println(*v.Name)
	}
}

