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
	"github.com/stianeikeland/go-rpio"
	"math/rand"
	"time"
)

type Configuration struct {
	Token string
	User  string
}

var titles []*string

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

func printTitle() {
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

func delete(i int) {
	titles[i] = titles[len(titles)-1] // Copy last element to index i.
	titles[len(titles)-1] = nil       // Erase last element (write zero value).
	titles = titles[:len(titles)-1]   // Truncate slice.
}

func merge(pin rpio.Pin) {
	idx := rand.Intn(len(titles))

	title := titles[idx]
	delete(idx)

	fmt.Printf("MERGED: %+s\n", *title)
	time.Sleep(1 * time.Second)
	listenToPin(pin)
}

func listenToPin(pin rpio.Pin) {


	for {
		if s := pin.Read(); s == rpio.High {
			//fmt.Println("HIGH")
			// We pushed the button so call merge!
			merge(pin)
		}

	}

}

//func listenToFinger() {
//	reader := bufio.NewReader(os.Stdin)
//	fmt.Print("Hit enter")
//
//
//	for {
//		reader.ReadString('\n')
//		break;
//	}
//
//	merge()
//}

func main() {
	rand.Seed(time.Now().Unix())
	ctx, client := getClient()
	//prs, err := getPullRequests(client, ctx)

	//if err != nil {
	//	fmt.Printf("\nerror: %v\n", err)
	//	return
	//}

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin := rpio.Pin(6)
	pin.Input()
	pin.PullDown() // Need to reduce the voltage because our push button will send 3.3v into pin

	printTitle()

	t, _ := getPullRequests(client, ctx)
	titles = t

	fmt.Println("\n-----------------------------------------------------------")
	color.Cyan("    Please press button to reduce Pull Request Backlog")
	fmt.Println("-----------------------------------------------------------")

	listenToPin(pin)

}

func getPullRequests(client *github.Client, ctx context.Context) ([]*string, error) {
	prs, _, err := client.PullRequests.List(ctx, "7shifts", "webapp", nil)

	titles := make([]*string, len(prs))

	for _, pr := range prs {
		titles = append(titles, pr.Title)
	}

	return titles, err
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
