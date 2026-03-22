package main

import (
	"fmt"
	"io"
	"net/http"
)

// middleware
func enableCORS(next http.Handler) http.Handler {
	// w = response object
	// r = incoming request
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// stop preflight request here
		if r.Method == http.MethodOptions {
			return
		}

		// OK middleware finished,now continue to the real route
		next.ServeHTTP(w, r)
	})
}

func main() {
	// mux = router
	// So mux decides: when a request comes in, which function should handle it?
	mux := http.NewServeMux()

	mux.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server!")
	})

	mux.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		// This reads the data sent from React.
		body, _ := io.ReadAll(r.Body)
		fmt.Println("Received:", string(body))
		fmt.Fprintln(w, "Status posted: "+string(body))
	})

	fmt.Println("Server running on http://localhost:8080")

	// wrap mux with CORS middleware
	// So flow becomes:enableCORS() runs first → then mux.HandleFunc("/get")
	http.ListenAndServe(":8080", enableCORS(mux))
}

// Flow:
// Incoming request
//   	↓
// CORS middleware
//   	↓
// Router (mux)
//   	↓
// Correct function (/get or /post)
//   	↓
// Response sent using w
