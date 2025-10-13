package protocol

type ListToolsRequest struct {
}

type ListToolsResult struct {
	Tools []*Tool `json:"tools"`
}

func NewListToolsResult() *ListToolsResult {
	return &ListToolsResult{
		Tools: make([]*Tool, 0),
	}
}

func (r *ListToolsResult) AddTool(tool *Tool) *ListToolsResult {
	r.Tools = append(r.Tools, tool)
	return r
}
