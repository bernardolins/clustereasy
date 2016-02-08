package templates

import (
	"github.com/bernardolins/clustereasy/scope"
	"log"
	"os"
	"text/template"
)

func ExecuteTemplate(temp string, scope scope.Scope) {
	t := template.New(scope.GetName())
	t, err := t.Parse(temp)

	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	t.Execute(os.Stdout, scope)
}
