package main

import (
    "fmt"
    "net/http"
    "sync"
)

var (
    // Текущее и следующее число Фибоначчи
    a, b = 0, 1
    // Счётчик запросов к основному эндпоинту
    requestCount int
    // Мьютекс для безопасного доступа из разных горутин
    mu sync.Mutex
)

func main() {
    // Основной эндпоинт — возвращает число Фибоначчи и увеличивает счётчик
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        // Возвращаем текущее число
        fmt.Fprint(w, a)
        // Переходим к следующему числу
        a, b = b, a+b
        // Увеличиваем счётчик запросов
        requestCount++
        mu.Unlock()
    })

    // Эндпоинт /metrics — возвращает количество запросов
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        mu.Lock()
        fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", requestCount)
        mu.Unlock()
    })

    // Запускаем сервер на порту 8080
    http.ListenAndServe(":8080", nil)
}