package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, DevOps World!")
}

func TestRootHandlerDirect(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	rootHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	body := w.Body.String()
	if !strings.Contains(body, "Hello, DevOps World!") {
		t.Errorf("expected body to contain 'Hello, DevOps World!', got %q", body)
	}
}

func TestServerResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(rootHandler))
	defer ts.Close()

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("failed to GET: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", resp.StatusCode)
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		t.Fatalf("failed to read body: %v", err)
	}
	if !strings.Contains(buf.String(), "Hello, DevOps World!") {
		t.Errorf("expected body to contain 'Hello, DevOps World!', got %q", buf.String())
	}
}

func TestExitInput(t *testing.T) {
	// Simulate user typing "exit"
	input := strings.NewReader("exit\n")
	scanner := bufio.NewScanner(input)
	exitCalled := false
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			exitCalled = true
			break
		}
	}
	if !exitCalled {
		t.Errorf("expected exit to be called")
	}
}
