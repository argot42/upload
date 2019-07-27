package utils

import (
	"testing"
)

func TestSplitFile(t *testing.T) {
	files := []string{"home.txt", "foo.tar.gz", ".bar.bin", "a.b", "onlyfile.", ".hiddenfile"}
	expected := [][]string{
		[]string{"home", ".txt"},
		[]string{"foo", ".tar.gz"},
		[]string{".bar", ".bin"},
		[]string{"a", ".b"},
		[]string{"onlyfile.", ""},
		[]string{".hiddenfile", ""},
	}

	for i, f := range files {
		filename, ext := SplitFile(f)
		if filename != expected[i][0] {
			t.Error(filename, "should be", expected[i][0])
		}
		if ext != expected[i][1] {
			t.Error(ext, "should be", expected[i][1])
		}
	}
}

func TestSplitPath(t *testing.T) {
	paths := []string{"/home/argot/file.txt", "/file.gz.tar", "./.hidden.zip", "foo/bar", "f/o/o/ba.r"}
	expected := [][]string{
		[]string{"/home/argot/", "file", ".txt"},
		[]string{"/", "file", ".gz.tar"},
		[]string{"./", ".hidden", ".zip"},
		[]string{"foo/", "bar", ""},
		[]string{"f/o/o/", "ba", ".r"},
	}

	for i, path := range paths {
		dir, f, ext := SplitPath(path)
		if dir != expected[i][0] {
			t.Error(dir, "should be", expected[i][0])
		}
		if f != expected[i][1] {
			t.Error(f, "should be", expected[i][1])
		}
		if ext != expected[i][2] {
			t.Error(ext, "should be", expected[i][2])
		}
	}
}
