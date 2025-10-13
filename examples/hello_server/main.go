package main

import (
	"context"
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
	server := gomcp.NewServer("hello", "", "1.0.0")

	server.AddTool(&gomcp.Tool{
		Name:        "hello",
		Title:       "Hello World",
		Description: "This is a tool that says hello world",
		InputSchema: protocol.NewInputSchema(),
		Handler: func(ctx context.Context, arguments *gomcp.ToolArguments) *protocol.CallToolsResult {
			return protocol.NewCallToolsResult().AddContent(protocol.NewTextContent().SetText("Hello world"))
		},
	})

	server.AddResource(&gomcp.Resource{
		Name: "test",
		Uri:  "gomcp://test",
		Handler: func(ctx context.Context) *protocol.ReadResourceResult {
			return protocol.NewReadResourceResult().AddContent(
				protocol.NewTextResourceContents("Hello world", "gomcp://test").SetMimeType("text/plain"),
			)
		},
	})

	server.AddTool(&gomcp.Tool{
		Name:        "add",
		Title:       "Add",
		Description: "Adds two numbers",
		InputSchema: protocol.NewInputSchema().
			SetProperty("a", protocol.NewNumberProperty("The first number")).
			SetProperty("b", protocol.NewNumberProperty("The second number")).
			SetRequired("a", "b"),
		OutputSchema: protocol.NewOutputSchema().
			SetProperty("result", protocol.NewNumberProperty("The result")),
		Handler: func(ctx context.Context, arguments *gomcp.ToolArguments) *protocol.CallToolsResult {
			a, err := arguments.Number("a")
			if err != nil {
				return protocol.NewCallToolsResult().AddContent(protocol.NewTextContent().SetText(err.Error())).SetIsError(true)
			}
			b, err := arguments.Number("b")
			if err != nil {
				return protocol.NewCallToolsResult().AddContent(protocol.NewTextContent().SetText(err.Error()))
			}

			return protocol.NewCallToolsResult().
				SetStructuredContent(map[string]any{
					"result": a + b,
				})
		},
	})

	transport := gomcp.NewHttpTransport(server)
	http.HandleFunc("/mcp", transport.Handle)
	slog.Info("Starting server", "addr", addr)
	http.ListenAndServe(addr, nil)
}
