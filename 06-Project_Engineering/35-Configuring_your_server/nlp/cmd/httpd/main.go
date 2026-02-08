package main

import (
	"encoding/json"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/dancoe77/nlp"
	"github.com/dancoe77/nlp/stemmer"
)

var (
	stemCalls = expvar.NewInt("stem.calls")
)

/*
configuration
- defaults < configuration file < environment variables < command line
-- configuration: YAML, TOML ... (not stdlib)
-- environment: os.Getenv
-- command line: flag

external
- viper + cobra
- ardanlabs/conf
*/

var config struct {
	Addr string
}

func main() {

	config.Addr = os.Getenv("NLP_ADDR")
	if config.Addr == "" {
		config.Addr = ":8080"
	}
	flag.StringVar(&config.Addr, "addr", config.Addr, "address to listen on")
	flag.Parse()

	//IMPORTANT: validate config
	if err := health(); err != nil {
		fmt.Fprintf(os.Stderr, "error: health check - %s\n", err)
		os.Exit(1)
	}

	api := API{
		log: slog.Default().With("app", "nlp"),
	}
	// Routing
	http.HandleFunc("GET /health", api.healthHandler)
	http.HandleFunc("POST /tokenize", api.tokenizeHandler)
	http.HandleFunc("GET /stem/{word}", api.stemHandler)

	api.log.Info("server starting", "address", config.Addr)
	if err := http.ListenAndServe(config.Addr, nil); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		if err != nil {
			return
		}
		os.Exit(1)
	}

}

type API struct {
	log *slog.Logger
}

func (a *API) stemHandler(w http.ResponseWriter, r *http.Request) {
	stemCalls.Add(1)
	word := r.PathValue("word")
	a.log.Info("stem", "word", word)
	fmt.Fprintln(w, stemmer.Stem(word))
}

func (a *API) tokenizeHandler(w http.ResponseWriter, r *http.Request) {
	// Step 1: Read data, parse & *validate*
	// TODO: http.MaxBytesReader
	data, err := io.ReadAll(r.Body)
	if err != nil {
		a.log.Error("read", "error", err, "remote", r.RemoteAddr)
		http.Error(w, "can't read", http.StatusBadRequest)
		return
	}
	text := string(data)
	if len(text) == 0 {
		a.log.Error("read", "error", "empty request")
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

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	if err := health(); err != nil {
		a.log.Error("health", "error", err)
		http.Error(w, "health check failed", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "OK")
}

func health() error {
	// TODO: Actual health check
	return nil
}
