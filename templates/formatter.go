package templates

import (
	"log"
	"os"
	"text/template"
)

func ExecuteTemplate(temp string, model interface{}) {
	t := template.New("template")
	t, err := t.Parse(temp)

	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	t.Execute(os.Stdout, model)
}
