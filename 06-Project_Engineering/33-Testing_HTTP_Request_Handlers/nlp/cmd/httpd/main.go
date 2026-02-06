package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dancoe77/nlp"
	"github.com/dancoe77/nlp/stemmer"
)

func main() {
	// Routing
	http.HandleFunc("GET /health", healthHandler)
	http.HandleFunc("POST /tokenize", tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", stemHandler)

	addr := ":8080"
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

}

func stemHandler(w http.ResponseWriter, r *http.Request) {
	word := r.PathValue("word")
	fmt.Fprintln(w, stemmer.Stem(word))
}

func tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read data, parse & *validate*
	// TODO: http.MaxBytesReader
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}
	text := string(data)
	if len(text) == 0 {
		http.Error(w, "empty request", http.StatusBadRequest)
		return
	}
	// Step 2: Work
	tokens := nlp.Tokenize(text)

	// Step 3: Encode response
	w.Header().Set("content-type", "application/json")
	resp := map[string]any{
		"tokens": tokens,
	}
	json.NewEncoder(w).Encode(resp)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		http.Error(w, "health check failed", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "OK")
}

func health() error {
	// TODO: Actual health check
	return nil
}
