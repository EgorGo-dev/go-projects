package main

import (
    "log/slog"
    "net/http"
    "os"
)

func Logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Логируем метод и путь запроса
        slog.Info("incoming request", "method", r.Method, "path", r.URL.Path)

        // Передаём управление следующему обработчику
        next.ServeHTTP(w, r)
    })
}

// helloHandler возвращает простой ответ.
func helloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, middleware!"))
}

func main() {
    // Настраиваем slog для вывода в консоль
    logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
    slog.SetDefault(logger)

    // Создаём роутер
    mux := http.NewServeMux()

    // Оборачиваем обработчик /hello в middleware Logger
    mux.Handle("/hello", Logger(http.HandlerFunc(helloHandler)))

    http.ListenAndServe(":8080", mux)
}