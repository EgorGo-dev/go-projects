package main

import (
	"fmt"
	"net/http"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EgorGo-dev")
}

func main() {
	http.HandleFunc("/", handle)
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", nil); err !=nil {
		log.Fatalf("Ошибка: %v", err)
	}
}