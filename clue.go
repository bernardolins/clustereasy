package main

import (
	"fmt"
	"github.com/bernardolins/clustereasy/cmd/clue"
	"os"
)

func main() {
	clue := cmd.NewClue()

	if err := clue.Command().Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
