package main

import (
	"context"
	"fmt"
	"net/http"
)

type contextKey string

const usernameKey contextKey = "username"

// Authorization — middleware для проверки Basic Auth.
func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		// ok == false, если заголовок отсутствует или имеет неверный формат
		if !ok || username == "" || password == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), usernameKey, username)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// answerHandler — обработчик для /answer/ (имя с маленькой буквы, как требует тест)
func answerHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value(usernameKey).(string)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Welcome, %s!", username)
}

func main() {
	http.HandleFunc("/answer/", Authorization(answerHandler))

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}