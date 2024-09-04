package main

import (
    "fmt"
    "net/http"
    "sync"
)

var mu sync.Mutex

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        mu.Lock() // Lock the mutex
        defer mu.Unlock() // Ensure the mutex is unlocked after handling the request
        fmt.Fprintf(w, "Hello, World!")
    })

    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
