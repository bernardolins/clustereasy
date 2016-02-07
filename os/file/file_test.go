package file

import (
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestLoad(t *testing.T) {
	data := "Hello, World!"
	expect := []byte(data)

	f := createFile("test-file", data, "", t)

	if got := Load(f.Name()); !equal(expect, got) {
		t.Fatalf("expect %q\ngot: %q", expect, got)
	}
}

func TestCreateFile(t *testing.T) {
	file := CreateFile("../../test/", "tempFile")

	expect := "tempFile"
	got := filename(file.Name())

	if expect != got {
		t.Fatalf("expect filename to be %s, but got %s", expect, got)
	}
}

func TestFilterFileByExtension(t *testing.T) {
	files := []string{
		"file1.yml",
		"file2.html",
		"file3.go",
		"file4.yml",
	}

	expect := []string{
		"file1.yml",
		"file4.yml",
	}

	got := FilterFilesByExtension(files, "yml")

	if !equal(expect, got) {
		t.Fatalf("expect filtered files to be %v, but got %v", expect, got)
	}

	expect = []string{}

	got = FilterFilesByExtension(files, "jpg")

	if !equal(expect, got) {
		t.Fatalf("expect filtered files to be %v, but got %v", expect, got)
	}
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

func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

func filename(path string) string {
	s := strings.Split(path, "/")
	size := len(s)

	return s[size-1]
}
