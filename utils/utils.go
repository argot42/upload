package utils

import "flag"

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
