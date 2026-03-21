package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// POST route FIRST
	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		body, _ := io.ReadAll(r.Body)
		fmt.Println("Received:", string(body))

		fmt.Fprintln(w, "Status posted: "+string(body))
	})

	// Root route AFTER
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server!")
	})

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
