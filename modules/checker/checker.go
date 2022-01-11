package checker

import (
	"net/http"
)

type Checker struct {
	apiUrl string
	apiVersion string
}

// Begin and do until the end check process
func (c Checker) start() {
	c.apiUrl = "https://discord.com/api"
	c.apiVersion = "v9"

	// CHECK IF PROXY

	// CREATE HTTP CLIENT
	client := &http.Client{}

	req, _ := http.NewRequest(
		"GET",
		c.apiUrl + "/" + c.apiVersion + "/users/@me",
	)

	// CREATE HEADERS

	// CHECK
}