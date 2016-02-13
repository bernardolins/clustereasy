package cmd

import (
	"fmt"
	"github.com/bernardolins/clustereasy/cloudconfig/header"
	"github.com/bernardolins/clustereasy/cloudconfig/keys/coreos"
	"github.com/bernardolins/clustereasy/os/dir"
	"github.com/bernardolins/clustereasy/os/file"
	//"github.com/bernardolins/clustereasy/scope"
	//"github.com/bernardolins/clustereasy/scope/coreos"
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

	header := header.New()

	for _, node := range initData.Cluster.Nodes {
		coreos := coreos.ConfigureCoreOS(node, initData.Cluster)
		checkUnitContentFile(&coreos, generate.units)
		//templates.ExecuteTemplate(templates.ScopeTemplateContent(), coreos)

		templates.ExecuteTemplate(templates.Header(), header)
		templates.ExecuteTemplate(templates.CoreOS(), coreos)
	}
}

func checkUnitContentFile(coreos *coreos.CoreOS, path string) {
	if path != "" {
		for _, unit := range coreos.Units() {
			if dir.HasFile(path, unit.Name()) {
				content := file.Load(path + unit.Name())
				unit.SetContent(string(content))
			}
		}
	}
}

func (generate *GenerateCommand) validateInput() error {
	if generate.input == "" {
		return fmt.Errorf("Input file path missing!\nTry something like 'cluectl generate -i /some/path/to/file.yml'")
	}

	return nil
}
