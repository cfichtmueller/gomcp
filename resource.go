package gomcp

import (
	"context"
	"errors"

	"github.com/cfichtmueller/gomcp/protocol"
)

type Resource struct {
	Name    string
	Uri     string
	Handler func(ctx context.Context) *protocol.ReadResourceResult
}

var ErrNoSuchResource = errors.New("no such resource")

type ResourceTemplate struct {
	Description string
	MimeType    string
	Name        string
	Title       string
	UriTemplate string
	// Read attempts to read a resource using the given URI. If the URI cannot be resolved using
	// this template, it returns ErrNoSuchResource.
	Read func(ctx context.Context, uri string) (*protocol.ReadResourceResult, error)
}
