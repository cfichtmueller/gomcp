package protocol

import "github.com/cfichtmueller/gomcp/schema"

type CallToolsRequest struct {
	Params CallToolsParams `json:"params"`
}

type CallToolsParams struct {
	Name      string   `json:"name"`
	Arguments schema.M `json:"arguments"`
}

type CallToolsResult struct {
	Meta              schema.M `json:"_meta,omitempty"`
	Content           schema.A `json:"content"`
	IsError           *bool    `json:"isError,omitempty"`
	StructuredContent schema.M `json:"structuredContent,omitempty"`
}

func NewCallToolsResult() *CallToolsResult {
	return &CallToolsResult{
		Content: make(schema.A, 0),
	}
}

func (r *CallToolsResult) AddContent(content any) *CallToolsResult {
	r.Content = append(r.Content, content)
	return r
}

func (r *CallToolsResult) SetStructuredContent(content map[string]any) *CallToolsResult {
	r.StructuredContent = content
	return r
}

func (r *CallToolsResult) SetIsError(isError bool) *CallToolsResult {
	r.IsError = &isError
	return r
}
