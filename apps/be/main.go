package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle requests to the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Set header so React can fetch without CORS issues (if needed)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, "Hello from Go server!")
	})

	// Start the server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
