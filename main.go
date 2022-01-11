package main

import (
	"fmt"
	"os"

	"github.com/Raskolnikov404/checker/modules"
)

var (
	tm modules.TokenManager
)

func main() {
	fmt.Println("Começando o Checker!")

	tm.Init()

	accounts, err := modules.ReadLines("files/accounts.txt")
	if err != nil {
		fmt.Println("Houve um erro!")
		os.Exit(1)
	}

	for _, account := range accounts {

	}
}
