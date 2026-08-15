package main 

import (
	"fmt"
	"net/http"
	"log"
)

func handle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Ok let`s Go")
}

func main() {
	http.HandleFunc("/", handle)
	fmt.Println("Сервер запущен на http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка при запуске сервера %v", err)
	}
}