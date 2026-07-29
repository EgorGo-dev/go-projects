package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// User — модель пользователя
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Store — потокобезопасное хранилище пользователей в памяти
type Store struct {
	mu    sync.Mutex
	users map[int]User
	nextID int
}

// NewStore — конструктор хранилища
func NewStore() *Store {
	return &Store{
		users:  make(map[int]User),
		nextID: 1,
	}
}

// CreateUser создаёт нового пользователя и возвращает его.
// Потокобезопасен: использует мьютекс.
func (s *Store) CreateUser(name string, age int) User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := User{
		ID:   s.nextID,
		Name: name,
		Age:  age,
	}
	s.users[s.nextID] = user
	s.nextID++
	return user
}

// GetUser возвращает пользователя по id и флаг успеха.
// Потокобезопасен.
func (s *Store) GetUser(id int) (User, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.users[id]
	return user, ok
}

// responseWriter — обёртка для http.ResponseWriter, перехватывающая статус ответа
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader перехватывает вызов и запоминает код статуса
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

// loggingMiddleware — middleware для логирования запросов
func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Оборачиваем оригинальный ResponseWriter, чтобы захватить статус
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)
		logger.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.statusCode),
			slog.Duration("duration", duration),
		)
	})
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	store := NewStore()

	mux := http.NewServeMux()

	// POST /users — создание пользователя
	mux.HandleFunc("POST /users", func(w http.ResponseWriter, r *http.Request) {
		var input struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "bad request")
			return
		}

		user := store.CreateUser(input.Name, input.Age)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	})

	// GET /users/{id} — получение пользователя
	mux.HandleFunc("GET /users/", func(w http.ResponseWriter, r *http.Request) {
		// Извлекаем id из пути: "/users/123" -> "123"
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		if idStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "bad request")
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, "bad request")
			return
		}

		user, ok := store.GetUser(id)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "not found")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)
	})

	// Оборачиваем mux в middleware
	handler := loggingMiddleware(logger, mux)

	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}