package main

import (
	"fmt"

	"github.com/Raskolnikov404/checker/modules/checker"
	"github.com/Raskolnikov404/checker/modules/tokenmanager"
)

var (
	dTokenManager tokenmanager.TokenManager
	dChecker      checker.Checker
)

func main() {
	fmt.Println("Começando o Checker!")

	dTokenManager = tokenmanager.TokenManager{}
	dTokenManager.Init()

	dChecker = checker.Checker{}
}
