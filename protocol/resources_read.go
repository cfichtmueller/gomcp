package protocol

type ReadResourceRequest struct {
	Params ReadResourceParams `json:"params"`
}

type ReadResourceParams struct {
	Uri string `json:"uri"`
}

type ReadResourceResult struct {
	Contents []any `json:"contents"`
}

func NewReadResourceResult() *ReadResourceResult {
	return &ReadResourceResult{
		Contents: make([]any, 0),
	}
}

func (r *ReadResourceResult) AddContent(content any) *ReadResourceResult {
	r.Contents = append(r.Contents, content)
	return r
}
