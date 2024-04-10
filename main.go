package main

import (
	"fmt"
	"net/http"
)

// main is the entry point of the program.
func main() {
	// Create a new HTTP server.
	server := &http.Server{
		Addr:    ":3000",                        // Listen on port 3000.
		Handler: http.HandlerFunc(basicHandler), // Set the handler function.
	}

	// Start the server and listen for incoming requests.
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Failed to Listen to Server:", err)
	}
}

// basicHandler is the handler function for the root route.
func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World")) // Write "Hello World" as the response.
}
