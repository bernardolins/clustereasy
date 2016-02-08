package main

import (
	"fmt"
	"github.com/bernardolins/clustereasy/cmd"
	"os"
)

func main() {
	if err := cmd.ClueCommand.Execute(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
