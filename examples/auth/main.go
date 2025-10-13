package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"slices"
	"strings"

	"github.com/cfichtmueller/gomcp"
	"github.com/cfichtmueller/gomcp/protocol"
)

type userKey struct{}

var UserKey = userKey{}

func main() {
	addr := os.Getenv("LISTEM_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := gomcp.NewServer()

	server.AddResource(&gomcp.Resource{
		Name: "user",
		Uri:  "gomcp://user",
		Handler: func(ctx context.Context) *protocol.ReadResourceResult {
			return protocol.NewReadResourceResult().AddContent(protocol.NewTextResourceContents(ctx.Value(UserKey).(string), "gomcp://user"))
		},
	})

	transport := gomcp.NewHttpTransport(server)
	http.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		ctx, ok := authenticate(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next := r.WithContext(ctx)
		transport.Handle(w, next)
	})
	slog.Info("Starting server", "addr", addr)
	http.ListenAndServe(addr, nil)
}

func authenticate(r *http.Request) (context.Context, bool) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		return nil, false
	}
	token := strings.TrimPrefix(auth, "Bearer ")
	if !slices.Contains([]string{"alice", "bob"}, token) {
		return nil, false
	}
	return context.WithValue(r.Context(), UserKey, token), true
}
