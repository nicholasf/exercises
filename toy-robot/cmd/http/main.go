package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/nicholasf/go-exercises/toy-robot/pkg/entities"
)

var logger *slog.Logger

func init() {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
}

func main() {
	err := runHTTP()

	entities.NewBoard()

	if err != nil {
		slog.Error("Unable to start server")
		os.Exit(1)
	}
}

func runHTTP() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /report", reporter)
	mux.HandleFunc("POST /place", placer)

	s := http.Server{
		Addr:         ":8080",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 90 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	return s.ListenAndServe()
}

// handles the REPORT command via a GET
// these functions will unpack the protocol and move to the usecases
func reporter(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Report"))
}

func placer(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Place"))
}
