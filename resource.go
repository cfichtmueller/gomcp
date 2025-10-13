package gomcp

import (
	"context"

	"github.com/cfichtmueller/gomcp/protocol"
)

type Resource struct {
	Name    string
	Uri     string
	Handler func(ctx context.Context) *protocol.ReadResourceResult
}

type ResourceResolver struct {
	List func(ctx context.Context) ([]*protocol.Resource, error)
	Read func(ctx context.Context, uri string) (*protocol.ReadResourceResult, error)
}
