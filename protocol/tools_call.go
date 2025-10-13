package protocol

type CallToolsRequest struct {
	Params CallToolsParams `json:"params"`
}

type CallToolsParams struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

type CallToolsResult struct {
	Meta              map[string]any `json:"_meta,omitempty"`
	Content           []any          `json:"content"`
	IsError           *bool          `json:"isError,omitempty"`
	StructuredContent map[string]any `json:"structuredContent,omitempty"`
}

func NewCallToolsResult() *CallToolsResult {
	return &CallToolsResult{
		Content: make([]any, 0),
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
