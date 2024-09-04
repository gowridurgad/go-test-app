package main

import (
    "fmt"
    "net/http"
)

func setupRoutes() {
    http.HandleFunc("/", helloHandler)
}

func startServer() {
    fmt.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Failed to start server:", err)
    }
}
