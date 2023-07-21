package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(rw, "Full Cycle Rocks!!")
    })

    log.Fatal(http.ListenAndServe(":8000", nil))
}