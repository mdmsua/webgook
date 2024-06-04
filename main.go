package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var infoLog = log.New(os.Stdout, "[INFO] ", log.LstdFlags)
var port = os.Getenv("PORT")

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	infoLog.Printf("RequestURI: %s, RemoteAddr: %s", r.RequestURI, r.RemoteAddr)
	w.WriteHeader(http.StatusNoContent)
}
