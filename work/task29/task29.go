package main

import (
    "fmt"
    "net/http"
    "regexp"
)


func onlyEnglishLetters(s string) bool {
    re := regexp.MustCompile(`^[a-zA-Z]+$`)
    return re.MatchString(s)
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Получаем параметр name из URL
        name := r.URL.Query().Get("name")

        // Если параметра нет или он пустой
        if name == "" {
            fmt.Fprint(w, "hello stranger")
            return
        }

        // Если имя содержит не только английские буквы
        if !onlyEnglishLetters(name) {
            fmt.Fprint(w, "hello dirty hacker")
            return
        }

        fmt.Fprintf(w, "hello %s", name)
    })

    http.ListenAndServe(":8080", nil)
}