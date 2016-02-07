package dir

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLs(t *testing.T) {
	dirname := createDir("test-dir", "", t)

	var files []string

	for i := 0; i < 1; i++ {
		f := createFile("test-file-", "test file", dirname, t)
		filename := filename(f.Name())
		files = append(files, filename)
	}

	if filenames := Ls(dirname); !equal(filenames, files) {
		t.Fatalf("expect %q\ngot: %q\n", filenames, files)
	}
}

func TestCreateDir(t *testing.T) {
	dirName := "../../test/tempDir"
	CreateDir(dirName)

	//contents := Ls("../../test/")

	os.Remove(dirName)

}

// -------- Helpers

func createFile(name, contents, dir string, t *testing.T) *os.File {
	f, err := ioutil.TempFile(dir, name)

	if err != nil {
		t.Fatal(err)
	}

	data := contents

	err = ioutil.WriteFile(f.Name(), []byte(data), 0644)

	if err != nil {
		t.Fatal(err)
	}

	return f
}

func createDir(name, path string, t *testing.T) string {
	dirname, err := ioutil.TempDir(path, name)

	if err != nil {
		t.Fatal(err)
	}

	return dirname
}

func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func filename(path string) string {
	s := strings.Split(path, "/")
	size := len(s)

	return s[size-1]
}
