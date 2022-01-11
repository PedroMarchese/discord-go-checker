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
	Tokens []account
}

// METHODS
func (tm TokenManager) Init() {
	tm.count = 0
	err := tm.readAccounts()
	if err != nil {
		panic("error on read accounts")
	}
}

func (tm TokenManager) readAccounts() error {
	file, err := ioutil.ReadFile("../files/accounts.txt")
	accounts := strings.Split(string(file), "\n")

	if len(accounts) < 1 {
		return errors.New("couldn't detect any account")
	}

	for _, account := range accounts {
		accArray := strings.Split(account, ":")
		tm.setTokenInfo(accArray[0], accArray[1], accArray[2])
	}

	return err
}

// GETTERS
// func (tm TokenManager) getTokensCount() int {
// 	return len(tm.Tokens)
// }

// SETTERS
func (tm TokenManager) setTokenInfo(token string, password string, email string) {
	newAccount := account{token, password, email}
	tm.Tokens = append(tm.Tokens, newAccount)
}
