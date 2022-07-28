package main

import (
	"github.com/ilijamt/vault-token-helper/cmd/vault-token-helper/command"
	"os"
)

func main() {
	var err = command.Execute()
	if err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
