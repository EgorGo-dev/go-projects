package main

import (
    "fmt"
    "net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
    // Получаем значение параметра msg из строки запроса
    msg := r.URL.Query().Get("msg")
    if msg == "" {
        fmt.Fprint(w, "empty")
        return
    }
    fmt.Fprint(w, msg)
}

func main() {
    http.HandleFunc("/echo", echoHandler)
    http.ListenAndServe(":8080", nil)
}