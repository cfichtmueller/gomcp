package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/cfichtmueller/gomcp"
)

func main() {
	addr := os.Getenv("LISTEM_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := gomcp.NewServer()

	transport := gomcp.NewHttpTransport(server)
	http.HandleFunc("/mcp", transport.Handle)
	slog.Info("Starting server", "addr", addr)
	http.ListenAndServe(addr, nil)
}
