package tokenmanager

import (
	"errors"
	"io/ioutil"
	"strings"
)

// STRUCTS HERE
type account struct {
	token    string
	password string
	email    string
}

type TokenManager struct {
	count  int
	Tokens map[int]account
}

// METHODS
func (tm TokenManager) Init() {
	tm.count = 0
	tm.readAccounts()
}

func (tm TokenManager) setTokenInfo(token string, password string, email string) {
	tm.Tokens[tm.count] = account{token, password, email}
	tm.count++
}

func (tm TokenManager) readAccounts() error {
	file, err := ioutil.ReadFile("../files/accounts.txt")
	accounts := strings.Split(string(file), "\n")

	if len(accounts) < 1 {
		return errors.New("couldn't detect any account")
	}

	for _, account := range accounts {
		accArray := strings.Split(account, ":")
		tm.SetTokenInfo(accArray[0], accArray[1], accArray[2])
	}

	return err
}
