package modules

import (
	"errors"
	"io/ioutil"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := ioutil.ReadFile(path)

	accounts := strings.Split(string(file), "\n")
	if len(accounts) < 1 {
		return nil, errors.New("No account detected!")
	}

	return accounts, err
}
