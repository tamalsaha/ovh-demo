package main

import (
	"fmt"
	"github.com/ovh/go-ovh/ovh"
)

// PartialMe holds the first name of the currently logged-in user.
// Visit https://api.ovh.com/console/#/me#GET for the full definition
type PartialMe struct {
	Firstname string `json:"firstname"`
}

// Instantiate an OVH client and get the firstname of the currently logged-in user.
// Visit https://api.ovh.com/createToken/index.cgi?GET=/me to get your credentials.
func main() {
	var me PartialMe

	client, _ := ovh.NewClient(
		"ovh-eu",
		"YOUR_APPLICATION_KEY",
		"YOUR_APPLICATION_SECRET",
		"YOUR_CONSUMER_KEY",
	)
	client.Get("/me", &me)
	fmt.Printf("Welcome %s!\n", me.Firstname)
}
