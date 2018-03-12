// Copyright 2015 The go-github AUTHORS. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The basicauth command demonstrates using the github.BasicAuthTransport,
// including handling two-factor authentication. This won't currently work for
// accounts that use SMS to receive one-time passwords.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/fatih/color"
	"golang.org/x/oauth2"
	"sync"
)

type Configuration struct {
	Token string
	User  string
}

func printDoge() {
	fmt.Println("░▄░░░░░░░░░░░░░░▄")
	fmt.Println("░░░░░░░░▌▒█░░░░░░░░░░░▄▀▒▌")
	fmt.Println("░░░░░░░░▌▒▒█░░░░░░░░▄▀▒▒▒▐")
	fmt.Println("░░░░░░░▐▄▀▒▒▀▀▀▀▄▄▄▀▒▒▒▒▒▐")
	fmt.Println("░░░░░▄▄▀▒░▒▒▒▒▒▒▒▒▒█▒▒▄█▒▐")
	fmt.Println("░░░▄▀▒▒▒░░░▒▒▒░░░▒▒▒▀██▀▒▌")
	fmt.Println("░░▐▒▒▒▄▄▒▒▒▒░░░▒▒▒▒▒▒▒▀▄▒▒▌")
	fmt.Println("░░▌░░▌█▀▒▒▒▒▒▄▀█▄▒▒▒▒▒▒▒█▒▐")
	fmt.Println("░▐░░░▒▒▒▒▒▒▒▒▌██▀▒▒░░░▒▒▒▀▄▌")
	fmt.Println("░▌░▒▄██▄▒▒▒▒▒▒▒▒▒░░░░░░▒▒▒▒▌")
	fmt.Println("▀▒▀▐▄█▄█▌▄░▀▒▒░░░░░░░░░░▒▒▒")
}

func printTitle()  {
	fmt.Print(`      /$$$$$$            /$$ /$$,
     /$$__  $$          |__/| $$
    | $$  \ $$  /$$$$$$  /$$| $$  /$$$$$$
    | $$$$$$$$ /$$__  $$| $$| $$ /$$__  $$
    | $$__  $$| $$  \ $$| $$| $$| $$$$$$$$
    | $$  | $$| $$  | $$| $$| $$| $$_____/
    | $$  | $$|  $$$$$$$| $$| $$|  $$$$$$$
    |__/  |__/ \____  $$|__/|__/ \_______/
               /$$  \ $$
              |  $$$$$$/
               \______/
     /$$      /$$
    | $$$    /$$$
    | $$$$  /$$$$  /$$$$$$   /$$$$$$  /$$$$$$   /$$$$$$
    | $$ $$/$$ $$ /$$__  $$ /$$__  $$/$$__  $$ /$$__  $$
    | $$  $$$| $$| $$$$$$$$| $$  \__/ $$  \ $$| $$$$$$$$
    | $$\  $ | $$| $$_____/| $$     | $$  | $$| $$_____/
    | $$ \/  | $$|  $$$$$$$| $$     |  $$$$$$$|  $$$$$$$
    |__/     |__/ \_______/|__/      \____  $$ \_______/
                                     /$$  \ $$
                                    |  $$$$$$/
                                     \______/`)
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
		fmt.Printf("\nerror: %v\n", err)
	}
	return configuration
}

func main() {
	//ctx, client := getClient()
	//prs, err := getPullRequests(client, ctx)

	//if err != nil {
	//	fmt.Printf("\nerror: %v\n", err)
	//	return
	//}

	printTitle()

	fmt.Println("\n-----------------------------------------------------------")
	color.Cyan("    Please press button to reduce Pull Request Backblog")
	fmt.Println("-----------------------------------------------------------")

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()

}

func getPullRequests(client *github.Client, ctx context.Context) ([]*github.PullRequest, error) {
	prs, _, err := client.PullRequests.List(ctx, "7shifts", "webapp", nil)

	return prs, err
}

func getClient() (context.Context, *github.Client) {
	config := readToken()
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: config.Token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return ctx, client
}
