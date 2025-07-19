package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Start HTTP server in the background
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, DevOps World!")
	})
	go func() {
		fmt.Println("Starting server on :8080")
		// Listen and serve in a goroutine
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait for user to type "exit" this allows you to stop the server gracefully
	fmt.Println("Server started on :8080. Type 'exit' and press Enter to stop.")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			break
		}
	}
	fmt.Println("Shutting down server...")
	if err := server.Close(); err != nil {
		fmt.Println("Error shutting down server:", err)
		os.Exit(70) // Exit with code 70 internal software error if this is not handled correctly
	} else {
		fmt.Println("Server stopped.")
	}
}
