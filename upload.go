package main

import (
	"fmt"
	"github.com/argot42/upload/utils"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func main() {
	config := utils.GetFlags()
	// serves the react client
	http.Handle(config.FileServerPath, http.FileServer(http.Dir(config.FileServerDir)))
	// handles the uploads
	http.HandleFunc(config.UploadPath, func(w http.ResponseWriter, r *http.Request) { upload(w, r, config.UploadDir) })

	log.Println("Listening on", config.Address)
	log.Fatal(http.ListenAndServe(config.Address, nil))
}

func upload(w http.ResponseWriter, r *http.Request, uploadDir string) {

	switch r.Method {
	case "GET":
		w.Write([]byte("GET THE FUCK OUT"))

	case "POST":
		reader, err := r.MultipartReader()
		if err != nil {
			fmt.Fprintln(os.Stderr, "error getting reader:", err)
			return
		}

		if err = saveFiles(reader, uploadDir); err != nil {
			fmt.Fprintln(os.Stderr, "error getting files:", err)
			return
		}
	}
}

func saveFiles(r *multipart.Reader, dir string) error {
	for {
		// get file reader
		part, err := r.NextPart()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}

		absoluteDir, err := filepath.Abs(dir)
		if err != nil {
			return err
		}
		filePath := filepath.Join(absoluteDir, part.FileName())

		log.Println("saving", part.FileName(), "at", absoluteDir)

		// create new file with non colliding filename
		file, err := createNewFile(filePath)
		if err != nil {
			return err
		}

		// copy partfile to a writer
		_, err = io.Copy(file, part)
		if err != nil {
			return err
		}

		// close part
		err = part.Close()
		if err != nil {
			return err
		}
		// close file
		err = file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func createNewFile(filename string) (newFile *os.File, err error) {
	if _, err = os.Stat(filename); os.IsNotExist(err) {
		newFile, err = os.Create(filename)
		if err != nil {
			return nil, err
		}
	} else {
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		newFile, err = os.Create(timestamp + "_" + filename)
		if err != nil {
			return nil, err
		}
	}
	return
}
