package main

import "fmt"
//import "github.com/google/go-querystring/query"
import "github.com/google/go-github/github"
import "golang.org/x/oauth2"
import "context"
import (
	"encoding/json"
	"os"
)

type Configuration struct {
	Token string
}

func main() {
	fmt.Printf("Hello, world.\n")
	//client := github.NewClient(nil)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: readToken()},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "knames", nil)

	if err != nil {
		// wut
	}
	if repos != nil {
		// wut
	}
}

func readToken() string {
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
	return configuration.Token
}

