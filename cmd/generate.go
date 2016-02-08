package cmd

import (
	"fmt"
	"github.com/bernardolins/clustereasy/os/file"
	"github.com/bernardolins/clustereasy/scope/coreos"
	"github.com/bernardolins/clustereasy/setup"
	"github.com/bernardolins/clustereasy/templates"
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
	output  string
	units   string
}

func New() *GenerateCommand {
	generate := new(GenerateCommand)

	generate.Command = &cobra.Command{
		Use: "generate",
		Run: func(cmd *cobra.Command, args []string) {
			if err := generate.validateInput(); err != nil {
				log.Fatalf("%v", err)
				os.Exit(1)
			} else {
				generate.run()
			}
		},
	}

	generate.Command.Flags().StringVarP(&generate.input, "input", "i", "", "Path to input file")
	generate.Command.Flags().StringVarP(&generate.output, "output", "o", "./", "Path to store output files")
	generate.Command.Flags().StringVarP(&generate.units, "units", "u", "", "Path to units")

	return generate
}

func (generate *GenerateCommand) run() {
	input := file.Load(generate.input)
	initData := setup.Apply(input)

	for _, node := range initData.Cluster.Nodes {
		coreos := coreos.CreateScope(node, initData.Cluster)

		templates.ExecuteTemplate(templates.ScopeTemplateContent(), *coreos)
	}
}

func (generate *GenerateCommand) validateInput() error {
	if generate.input == "" {
		return fmt.Errorf("Input file path missing!\nTry something like 'cluectl generate -i /some/path/to/file.yml'")
	}

	return nil
}
