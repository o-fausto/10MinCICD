package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := &http.Server{Addr: ":8080"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, DevOps World!")
	})
	go func() {
		fmt.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("Server error:", err)
		}
	}()

	// Wait for termination signal (Ctrl+C, SIGTERM, etc.)
	fmt.Println("Server started on :8080. Waiting for termination signal (Ctrl+C or SIGTERM).")
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("Shutting down server...")
	if err := server.Close(); err != nil {
		fmt.Println("Error shutting down server:", err)
		os.Exit(70) // Exit with code 70 internal software error if this is not handled correctly
	} else {
		fmt.Println("Server stopped.")
	}
}
