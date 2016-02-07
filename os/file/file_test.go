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

	f := createTempFile("test-file", data, "", t)

	if got := Load(f.Name()); !equal(expect, got) {
		t.Fatalf("expect %q\ngot: %q", expect, got)
	}
}

func TestCreateFile(t *testing.T) {
	fileName := "tempFile"
	testDir := "../../test/"
	file := CreateFile(testDir, fileName)

	expect := "tempFile"
	got := filename(file.Name())

	if expect != got {
		t.Fatalf("expect filename to be %s\ngot %s", expect, got)
	}

	os.Remove(testDir + fileName)
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
		t.Fatalf("expect filtered files to be %v\ngot %v", expect, got)
	}

	expect = []string{}

	got = FilterFilesByExtension(files, "jpg")

	if !equal(expect, got) {
		t.Fatalf("expect filtered files to be %v\ngot %v", expect, got)
	}
}

// -------- Helpers

func createTempFile(name, contents, dir string, t *testing.T) *os.File {
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
