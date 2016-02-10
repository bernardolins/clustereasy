package dir

import (
	"io/ioutil"
	"log"
	"os"
)

func Ls(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	var filenames []string

	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	for _, f := range files {
		filenames = append(filenames, f.Name())
	}

	return filenames
}

func CreateDir(dirname string) {
	err := os.Mkdir(dirname, os.ModePerm)

	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}
}

func HasFile(dirname, filename string) bool {
	files := Ls(dirname)

	for _, file := range files {
		if file == filename {
			return true
		}
	}

	return false
}
