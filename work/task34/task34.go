package main

import (
	"context"
	"net/http"
	"regexp"
)

type contextKey string

const nameKey contextKey = "name"

var validName = regexp.MustCompile(`^[a-zA-Z]+$`)

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"
		}
		ctx := context.WithValue(r.Context(), nameKey, name)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name, _ := r.Context().Value(nameKey).(string)
		if name != "stranger" && !validName.MatchString(name) {
			name = "dirty hacker"
			ctx := context.WithValue(r.Context(), nameKey, name)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name, _ := r.Context().Value(nameKey).(string)
	w.Write([]byte("hello " + name))
}

func main() {
	handler := Sanitize(SetDefaultName(HelloHandler))
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}