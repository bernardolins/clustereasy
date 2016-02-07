package file

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Load(path string) []byte {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("error: %v", err)
		os.Exit(1)
	}

	return file
}

func CreateFile(path, filename string) *os.File {
	file, err := os.Create(path + filename)

	if err != nil {
		log.Fatalf("%v", err)
		os.Exit(1)
	}

	return file
}

func FilterFilesByExtension(fileNames []string, extension string) []string {
	filesWithExtension := []string{}

	for _, name := range fileNames {
		if strings.HasSuffix(name, extension) {
			filesWithExtension = append(filesWithExtension, name)
		}
	}

	return filesWithExtension
}
