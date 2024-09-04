package main

import (
    "fmt"
    "net/http"
    "sync"
)

var mu sync.Mutex

func helloHandler(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    fmt.Fprintf(w, "Hello, World!")
}
