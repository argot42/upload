package main

import (
    "log"
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir("./public")))
    http.HandleFunc("/upload", upload)

    log.Println("Listening on :1234")
    log.Fatal(http.ListenAndServe(":1234", nil))
}

func upload(w http.ResponseWriter, r *http.Request) {
    log.Println("method:", r.Method)

    switch(r.Method) {
    case "GET":
        w.Write([]byte("GET THE FUCK OUT"))

    case "POST":
        w.Write([]byte("xd"))
    }
}
