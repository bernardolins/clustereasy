package clue

import (
	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use: "clue",
	Run: func(cmd *cobra.Command, args []string) {
		printHelp()
	},
}

func printHelp() {

}
