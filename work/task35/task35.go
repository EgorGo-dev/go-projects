package main

import (
	"fmt"
	"net/http"
	"sync"
)

var (
	mu         sync.Mutex
	fibPrev    int64 = 0
	fibCurr    int64 = 1
	requestCnt int64
)

func Metrics(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			mu.Lock()
			requestCnt++
			mu.Unlock()
		}
		next(w, r)
	}
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	fmt.Fprint(w, fibPrev)

	fibPrev, fibCurr = fibCurr, fibPrev+fibCurr
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	cnt := requestCnt
	mu.Unlock()

	fmt.Fprintf(w, "rpc_duration_milliseconds_count %d", cnt)
}

func main() {
	http.HandleFunc("/", Metrics(FibHandler))
	http.HandleFunc("/metrics", Metrics(MetricsHandler))

	http.ListenAndServe(":8080", nil)
}