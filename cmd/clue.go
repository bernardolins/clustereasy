package cmd

import (
	"github.com/spf13/cobra"
)

var ClueCommand = &cobra.Command{
	Use: "clue",
	Run: func(cmd *cobra.Command, args []string) {
		printHelp()
	},
}

func printHelp() {
}
