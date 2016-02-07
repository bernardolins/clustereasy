package cmd

import (
	"github.com/spf13/cobra"
)

type Clue struct {
	command *cobra.Command
}

func NewClue() *Clue {
	clue := new(Clue)

	clue.command = &cobra.Command{
		Use: "clue",
		Run: func(cmd *cobra.Command, args []string) {
			clue.printHelp()
		},
	}

	return clue
}

func (clue *Clue) Command() *cobra.Command {
	return clue.command
}

func (clue *Clue) printHelp() {

}
