package main

import (
	"api/scripts"
	"fmt"
	"net/http"
	"os"
	"time"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(

		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			fmt.Printf("Origin: %s \n[%s] %s %s\n", r.RemoteAddr, time.Now().Format("2006-01-02 15:04:05"), r.Method, r.RequestURI)

			next.ServeHTTP(w, r)

			duration := time.Since(start)
			fmt.Printf("  └─ Completed in %v\n", duration)

		},
	)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/sequential", scripts.SequentialManager)
	mux.HandleFunc("/concurrent", scripts.ConcurrentManager)

	handler := corsMiddleware(loggerMiddleware(mux))

	port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

	fmt.Println("Server running on :", port)
	http.ListenAndServe(":"+port, handler)
}
