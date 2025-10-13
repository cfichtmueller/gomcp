package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"regexp"

	"github.com/cfichtmueller/gomcp"
	"github.com/cfichtmueller/gomcp/protocol"
)

var (
	postsRegex = regexp.MustCompile(`^gomcp://posts/([a-zA-Z0-9]+)$`)
	pagesRegex = regexp.MustCompile(`^gomcp://pages/([a-zA-Z0-9]+)$`)
)

func main() {
	addr := os.Getenv("LISTEM_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	server := gomcp.NewServer()

	server.AddResourceTemplate(&gomcp.ResourceTemplate{
		Name:        "posts",
		UriTemplate: "gomcp://posts/{id}",
		Read: func(ctx context.Context, uri string) (*protocol.ReadResourceResult, error) {
			matches := postsRegex.FindStringSubmatch(uri)
			if matches == nil {
				return nil, gomcp.ErrNoSuchResource
			}
			id := matches[1]
			content := fmt.Sprintf("Post %s", id)
			return protocol.NewReadResourceResult().AddContent(protocol.NewTextResourceContents(content, uri)), nil
		},
	})

	server.AddResourceTemplate(&gomcp.ResourceTemplate{
		Name:        "pages",
		UriTemplate: "gomcp://pages/{id}",
		Read: func(ctx context.Context, uri string) (*protocol.ReadResourceResult, error) {
			matches := pagesRegex.FindStringSubmatch(uri)
			if matches == nil {
				return nil, gomcp.ErrNoSuchResource
			}
			id := matches[1]
			content := fmt.Sprintf("Page %s", id)
			return protocol.NewReadResourceResult().AddContent(protocol.NewTextResourceContents(content, uri)), nil
		},
	})

	transport := gomcp.NewHttpTransport(server)
	http.HandleFunc("/mcp", transport.Handle)
	slog.Info("Starting server", "addr", addr)
	http.ListenAndServe(addr, nil)
}
