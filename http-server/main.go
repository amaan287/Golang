package main

import (
    "net/http"
    "fmt"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })
    
    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))
    
    fmt.Println("Server starting on port 80...")
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}
