package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
    http.HandleFunc("/", handleRoot)
    http.HandleFunc("/time", handleTime)

    fmt.Println("Server is starting on :8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the Go HTTP server!")
}

func handleTime(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format(time.RFC3339)
    fmt.Fprintf(w, "The current time is: %s", currentTime)
}