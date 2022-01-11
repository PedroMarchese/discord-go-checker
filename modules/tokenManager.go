package modules

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// STRUCTS HERE
type account struct {
	token    string
	password string
	email    string
}

type TokenManager struct {
	lastIndex int
	Tokens    map[int]account
}

// METHODS
func readLines() ([]string, error) {
	file, err := ioutil.ReadFile("../files/accounts.txt")

	accounts := strings.Split(string(file), "\n")
	if len(accounts) < 1 {
		fmt.Println("No account detected!")
		os.Exit(1)
	}

	return accounts, err
}

func (self TokenManager) Init() {
	self.lastIndex = 0
}

func (self TokenManager) SetTokenInfo(token string, password string, email string) {
	self.Tokens[self.lastIndex] = account{token, password, email}
	self.lastIndex++
}
