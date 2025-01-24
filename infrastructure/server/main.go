package server

import (
	"fmt"
	"io"
	"net/http"
)

func Server() {
    http.HandleFunc("/", getRoot)
    err := http.ListenAndServe(":3333", nil)
    if err != nil {
        fmt.Printf("Server error: Error to run server")
    }
}

func getRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("got request")
    io.WriteString(w, "Hello! this is root")
}
