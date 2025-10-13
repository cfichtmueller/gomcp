package protocol

// ListResourcesRequest is sent from the client to request a list of resources the server has.
type ListResourcesRequest struct {
}

// ListResourcesResult is server’s response to a resources/list request from the client.
type ListResourcesResult struct {
	Resources []*Resource `json:"resources"`
}

// NewListResourcesResult creates a new ListResourcesResult with an empty list of resources.
func NewListResourcesResult() *ListResourcesResult {
	return &ListResourcesResult{
		Resources: make([]*Resource, 0),
	}
}

func (r *ListResourcesResult) AddResource(resource *Resource) *ListResourcesResult {
	r.Resources = append(r.Resources, resource)
	return r
}

type ListResourcesTemplatesResult struct {
	ResourceTemplates []*ResourceTemplate `json:"resourceTemplates"`
}

func NewListResourcesTemplatesResult() *ListResourcesTemplatesResult {
	return &ListResourcesTemplatesResult{
		ResourceTemplates: make([]*ResourceTemplate, 0),
	}
}

func (r *ListResourcesTemplatesResult) AddResourceTemplate(resourceTemplate *ResourceTemplate) *ListResourcesTemplatesResult {
	r.ResourceTemplates = append(r.ResourceTemplates, resourceTemplate)
	return r
}
