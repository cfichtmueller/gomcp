package protocol

import "github.com/cfichtmueller/gomcp/schema"

type ReadResourceRequest struct {
	Params ReadResourceParams `json:"params"`
}

type ReadResourceParams struct {
	Uri string `json:"uri"`
}

type ReadResourceResult struct {
	Contents schema.A `json:"contents"`
}

func NewReadResourceResult() *ReadResourceResult {
	return &ReadResourceResult{
		Contents: make(schema.A, 0),
	}
}

func (r *ReadResourceResult) AddContent(content any) *ReadResourceResult {
	r.Contents = append(r.Contents, content)
	return r
}
