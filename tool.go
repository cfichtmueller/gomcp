package gomcp

import (
	"context"
	"fmt"

	"github.com/cfichtmueller/gomcp/protocol"
	"github.com/cfichtmueller/gomcp/schema"
)

type Tool struct {
	Name         string
	Title        string
	Description  string
	InputSchema  *protocol.InputSchema
	OutputSchema *protocol.OutputSchema
	Handler      func(ctx context.Context, arguments *ToolArguments) *protocol.CallToolsResult
}

func (t *Tool) Call(ctx context.Context, arguments *ToolArguments) *protocol.CallToolsResult {
	if t.Handler == nil {
		panic("tool handler is not set")
	}
	return t.Handler(ctx, arguments)
}

type ToolArguments struct {
	backing schema.M
}

func NewToolArguments(arguments schema.M) *ToolArguments {
	return &ToolArguments{
		backing: arguments,
	}
}

func (t *ToolArguments) Number(key string) (float64, error) {
	raw, ok := t.backing[key]
	if !ok {
		return 0, fmt.Errorf("key %s not found", key)
	}
	value, ok := raw.(float64)
	if !ok {
		return 0, fmt.Errorf("key %s is not a number", key)
	}
	return value, nil
}

func (t *ToolArguments) String(key string) (string, error) {
	raw, ok := t.backing[key]
	if !ok {
		return "", fmt.Errorf("key %s not found", key)
	}
	value, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("key %s is not a string", key)
	}
	return value, nil
}
