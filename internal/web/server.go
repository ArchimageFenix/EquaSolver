// File: server.go
// Purpose: configure the routes and start the HTTP server.
// Receives: the listening address from cmd/web/main.go.
// Previous stage: cmd/web/main.go.
// Next stage: handlers.go (one handler per route).
// Restrictions: no domain logic and no HTML; only wiring and server settings.

package web

import (
	"net/http"
	"time"
)

// Serve starts the server and blocks until it fails.
func Serve(address string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/solve", solveHandler)
	// Serve static Web resources directly from the embedded filesystem.
	// The files remain part of the EquaSolver binary at runtime.
	mux.Handle("/static/", http.FileServer(http.FS(webFiles)))
	server := &http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	return server.ListenAndServe()
}
