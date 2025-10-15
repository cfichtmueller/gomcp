package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/cfichtmueller/gomcp"
	"github.com/cfichtmueller/gomcp/protocol"
)

func main() {
	server := gomcp.NewServer("my-server", "", "1.0.0")

	server.AddTool(&gomcp.Tool{
		Name:        "hello",
		Title:       "Hello World",
		Description: "This is a tool that says hello world",
		InputSchema: protocol.NewInputSchema().
			SetProperty("name", protocol.NewStringProperty("The name to say hello to")).
			SetRequired("name"),
		Handler: func(ctx context.Context, arguments *gomcp.ToolArguments) *protocol.CallToolsResult {
			name, err := arguments.String("name")
			if err != nil {
				return protocol.NewCallToolsResult().AddContent(protocol.NewTextContent().SetText(err.Error())).SetIsError(true)
			}
			content := protocol.NewTextContent().SetText(fmt.Sprintf("Hello, %s!", name))
			return protocol.NewCallToolsResult().AddContent(content)
		},
	})

	transport := gomcp.NewHttpTransport(server)
	http.HandleFunc("/mcp", transport.Handle)
	slog.Info("Starting MCP server on 127.0.0.1:8080")
	http.ListenAndServe("127.0.0.1:8080", nil)
}
