package main

import (
    "fmt"
    "net/http"
)

func main() {
    a, b := 0, 1

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, a)

        a, b = b, a+b
    }) 

    http.ListenAndServe(":8080", nil)
}