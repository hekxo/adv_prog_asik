package main

import (
	"net/http"
)

func main() {
	// Set up a handler function for the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	// Start the server on port 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
