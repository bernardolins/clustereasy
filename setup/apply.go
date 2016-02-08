package setup

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

func Apply(data []byte) InitData {
	init := InitData{}

	err := yaml.Unmarshal(data, &init)

	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	return init
}
