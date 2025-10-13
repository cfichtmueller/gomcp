package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/cfichtmueller/gomcp"
	"github.com/cfichtmueller/gomcp/protocol"
)

func main() {
	addr := os.Getenv("LISTEM_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}

	server := gomcp.NewServer()

	server.AddResourceResolver(&gomcp.ResourceResolver{
		List: ListResources,
		Read: ReadResource,
	})

	transport := gomcp.NewHttpTransport(server)
	http.HandleFunc("/mcp", transport.Handle)
	slog.Info("Starting server", "addr", addr)
	http.ListenAndServe(addr, nil)
}

type Resource struct {
	Name    string
	Uri     string
	Content string
}

var resources = []Resource{
	{Name: "dyn1", Uri: "gomcp://dyn1", Content: "Dynamic resource 1"},
	{Name: "dyn2", Uri: "gomcp://dyn2", Content: "Dynamic resource 2"},
	{Name: "dyn3", Uri: "gomcp://dyn3", Content: "Dynamic resource 3"},
}

func resolveResource(uri string) *Resource {
	for _, r := range resources {
		if r.Uri == uri {
			return &r
		}
	}
	return nil
}

func ListResources(ctx context.Context) ([]*protocol.Resource, error) {
	res := make([]*protocol.Resource, len(resources))
	for i, r := range resources {
		res[i] = &protocol.Resource{
			Name: r.Name,
			Uri:  r.Uri,
		}
	}
	return res, nil
}

func ReadResource(ctx context.Context, uri string) (*protocol.ReadResourceResult, error) {
	r := resolveResource(uri)
	if r == nil {
		return nil, fmt.Errorf("resource not found")
	}
	return protocol.NewReadResourceResult().AddContent(protocol.NewTextResourceContents(r.Content, uri)), nil
}
