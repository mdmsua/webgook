package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var infoLog = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
var errorLog = log.New(os.Stderr, "[ERR] ", log.LstdFlags)
var port = os.Getenv("PORT")

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	infoLog.Printf("%s %s", r.Method, r.RequestURI)
	data, err := io.ReadAll(r.Body)
	if err != nil {
		errorLog.Println(err)
	} else {
		infoLog.Println(string(data))
	}
	w.WriteHeader(http.StatusNoContent)
}
