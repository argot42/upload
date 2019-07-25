package main

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestCopyPartFile(t *testing.T) {
	filename0 := "foo.txt"
	filename1 := "bar"

	file0, err := createNewFile(filename0)
	if err != nil {
		t.Fatal(err)
	}
	path0 := file0.Name()
	defer file0.Close()
	defer os.Remove(path0)

	file1, err := createNewFile(filename1)
	if err != nil {
		t.Fatal(err)
	}
	path1 := file1.Name()
	defer file1.Close()
	defer os.Remove(path1)

	file2, err := createNewFile(filename0)
	if err != nil {
		t.Fatal(err)
	}
	path2 := file2.Name()
	defer file2.Close()
	defer os.Remove(path2)

	file3, err := createNewFile(filename1)
	if err != nil {
		t.Fatal(err)
	}
	path3 := file3.Name()
	defer file3.Close()
	defer os.Remove(path3)

	if filepath.Base(file0.Name()) != filename0 {
		t.Error(file0.Name() + " should have been " + filename0)
	}
	if filepath.Base(file1.Name()) != filename1 {
		t.Error(file1.Name() + " should have been " + filename1)
	}
	if filepath.Base(file2.Name()) == filename0 {
		t.Error(file2.Name() + " should have been different to " + filename0)
	}
	if filepath.Base(file3.Name()) == filename1 {
		t.Error(file3.Name() + " should have been different to " + filename1)
	}
}

func TestSaveFiles(t *testing.T) {
	s := "--foo\r\nContent-Disposition: form-data; filename=\"bar.txt\"\r\nContent-Type: text/plain\r\n\r\nbar0\r\n" +
		"--foo\r\nContent-Disposition: form-data; filename=\"file.txt\"\r\nContent-Type: text/plain\r\n\r\nbar1\r\n" +
		"--foo--"

	reader := strings.NewReader(s)
	multipartReader := multipart.NewReader(reader, "foo")

	// check that filenames in the formdata are not taken
	for {
		part, err := multipartReader.NextPart()
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			break
		}

		if _, err = os.Stat(part.FileName()); !os.IsNotExist(err) {
			t.Fatal("file " + part.FileName() + " exists on directory")
		}
	}

	reader = strings.NewReader(s)
	multipartReader = multipart.NewReader(reader, "foo")
	err := saveFiles(multipartReader)
	if err != nil {
		t.Fatal(err)
	}

	reader = strings.NewReader(s)
	multipartReader = multipart.NewReader(reader, "foo")
	// check if files info is copy correctly
	for {
		part, err := multipartReader.NextPart()
		if err != nil {
			if err != io.EOF {
				t.Fatal(err)
			}
			break
		}

		file, err := os.Open(part.FileName())
		if err != nil {
			t.Fatal(err)
		}
		defer file.Close()
		defer os.Remove(part.FileName())
		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			t.Fatal(err)
		}
		partContent, err := ioutil.ReadAll(part)
		if err != nil {
			t.Fatal(err)
		}

		fileContentStr := string(fileContent)
		partContentStr := string(partContent)
		if fileContentStr != partContentStr {
			t.Error(fileContentStr + " should have been " + partContentStr)
		}
	}
}
