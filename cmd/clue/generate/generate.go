package generate

import (
	"github.com/bernardolins/clustereasy/cmd/clue"
	"github.com/spf13/cobra"
)

func init() {
	generateCommand := New()
	clue.Command.AddCommand(generateCommand.command)
}

type Command struct {
	command *cobra.Command
	units   string
}

func New() *Command {
	generate := new(Command)

	generate.command = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			generate.run()
		},
	}

	generate.command.Flags().StringVarP(&generate.units, "units", "u", "", "Path to units")

	return generate
}

func (generate *Command) run() {
}
