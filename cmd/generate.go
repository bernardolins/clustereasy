package cmd

import (
	"fmt"
	"github.com/bernardolins/clustereasy/os/file"
	"github.com/bernardolins/clustereasy/setup"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	generateCommand := New()
	ClueCommand.AddCommand(generateCommand.Command)
}

type GenerateCommand struct {
	Command *cobra.Command
	input   string
	units   string
}

func New() *GenerateCommand {
	generate := new(GenerateCommand)

	generate.Command = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generate.validateInput(); err != nil {
				generate.run()
			} else {
				log.Fatalf("%v", err)
				os.Exit(1)
			}
		},
	}

	generate.Command.Flags().StringVarP(&generate.input, "input", "i", "", "Path to input file")
	generate.Command.Flags().StringVarP(&generate.units, "units", "u", "", "Path to units")

	return generate
}

func (generate *GenerateCommand) run() {
	input := file.Load(generate.input)
	initData := setup.Apply(input)

	fmt.Printf("%+v", initData)
}

func (generate *GenerateCommand) validateInput() error {
	if generate.input == "" {
		return fmt.Errorf("Input file path missing!")
	}

	return nil
}
