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
