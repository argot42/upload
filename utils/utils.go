package utils

import (
	"flag"
	"path/filepath"
)

type Config struct {
	FileServerPath string
	FileServerDir  string
	UploadPath     string
	UploadDir      string
	Address        string
}

func GetFlags() Config {
	config := Config{
		*flag.String("FileServerPath", "/", "Path for the client"),
		*flag.String("FileServerDir", "app", "Directory path for the client"),
		*flag.String("UploadPath", "/upload", "Path for the upload handler"),
		*flag.String("UploadDir", "files", "Directory path for uploaded files"),
		*flag.String("Address", ":1234", "<host>:<port>"),
	}
	flag.Parse()
	return config
}

func SplitPath(path string) (directory string, filename string, extension string) {
	directory, filenameWithExt := filepath.Split(path)
	filename, extension = SplitFile(filenameWithExt)
	return
}

func SplitFile(file string) (filename string, extension string) {
	var i int

	for i = 0; i < len(file); i++ {
		if file[i] == '.' && i > 0 {
			break
		}
	}
	if i > 0 && i < len(file)-1 {
		filename = file[:i]
		extension = file[i:]
	} else {
		filename = file
	}

	return
}
