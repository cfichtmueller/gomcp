package protocol

type Annotations struct {
	Audience     []Role `json:"audience,omitempty"`
	LastModified string `json:"lastModified,omitempty"`
	Priority     *int   `json:"priority,omitempty"`
}

func (a *Annotations) SetPriority(priority int) {
	a.Priority = &priority
}

type ContentBlock struct {
	Meta        map[string]any `json:"_meta,omitempty"`
	Annotations *Annotations   `json:"annotations,omitempty"`
	Type        string         `json:"type"`
}

type TextContent struct {
	ContentBlock
	Text string `json:"text"`
}

func NewTextContent() *TextContent {
	return &TextContent{
		ContentBlock: ContentBlock{
			Type: "text",
		},
	}
}

func (t *TextContent) SetText(text string) *TextContent {
	t.Text = text
	return t
}

type Role string

var (
	RoleUser      Role = "user"
	RoleAssistant Role = "assistant"
)

type ResourceLink struct {
	ContentBlock
	Description string `json:"description,omitempty"`
	Name        string `json:"name"`
}

func NewResourceLink() *ResourceLink {
	return &ResourceLink{
		ContentBlock: ContentBlock{
			Type: "resource_link",
		},
	}
}

type Tool struct {
	Description  string        `json:"description,omitempty"`
	InputSchema  *InputSchema  `json:"inputSchema"`
	Name         string        `json:"name"`
	OutputSchema *OutputSchema `json:"outputSchema,omitempty"`
	Title        string        `json:"title,omitempty"`
}

func NewTool(name string) *Tool {
	return &Tool{
		InputSchema: NewInputSchema(),
		Name:        name,
	}
}

type InputSchema struct {
	Properties map[string]any `json:"properties,omitempty"`
	Required   []string       `json:"required,omitempty"`
	Type       string         `json:"type"`
}

func NewInputSchema() *InputSchema {
	return &InputSchema{
		Type: "object",
	}
}

func (i *InputSchema) SetProperty(key string, value any) *InputSchema {
	if i.Properties == nil {
		i.Properties = make(map[string]any)
	}
	i.Properties[key] = value
	return i
}

func (i *InputSchema) SetRequired(required ...string) *InputSchema {
	i.Required = required
	return i
}

type OutputSchema struct {
	Properties map[string]any `json:"properties,omitempty"`
	Required   []string       `json:"required,omitempty"`
	Type       string         `json:"type"`
}

func NewOutputSchema() *OutputSchema {
	return &OutputSchema{
		Type: "object",
	}
}

func (i *OutputSchema) SetProperty(key string, value any) *OutputSchema {
	if i.Properties == nil {
		i.Properties = make(map[string]any)
	}
	i.Properties[key] = value
	return i
}

func (i *OutputSchema) SetRequired(required ...string) *OutputSchema {
	i.Required = required
	return i
}

func NewStringProperty(description string) any {
	return map[string]any{
		"type":        "string",
		"description": description,
	}
}

func NewNumberProperty(description string) any {
	return map[string]any{
		"type":        "number",
		"description": description,
	}
}

// Resource is a known resource that the server is capable of reading.
type Resource struct {
	// A description of what this resource represents.
	//
	// This can be used by clients to improve the LLM’s understanding of available resources. It can be thought of like a “hint” to the model.
	Description string `json:"description,omitempty"`
	// The MIME type of this resource, if known.
	MimeType string `json:"mimeType,omitempty"`
	// Intended for programmatic or logical use, but used as a display name in past specs or fallback (if title isn’t present).
	Name string `json:"name"`
	// The size of the raw resource content, in bytes (i.e., before base64 encoding or any tokenization), if known.
	//
	// This can be used by Hosts to display file sizes and estimate context window usage.
	Size *float64 `json:"size,omitempty"`
	// Intended for UI and end-user contexts — optimized to be human-readable and easily understood, even by those unfamiliar with domain-specific terminology.
	//
	// If not provided, the name should be used for display (except for Tool, where annotations.title should be given precedence over using name, if present).
	Title string `json:"title,omitempty"`
	// The URI of this resource.
	Uri string `json:"uri"`
}

// NewResource creates a new Resource instance with the given name and URI.
// This is a convenience constructor for creating basic resource objects.
func NewResource(name, uri string) *Resource {
	return &Resource{
		Name: name,
		Uri:  uri,
	}
}

type TextResourceContents struct {
	MimeType string `json:"mimeType,omitempty"`
	Text     string `json:"text"`
	Uri      string `json:"uri"`
}

func NewTextResourceContents(text, uri string) *TextResourceContents {
	if text == "" {
		panic("text is not set")
	}
	if uri == "" {
		panic("uri is not set")
	}
	return &TextResourceContents{
		Text: text,
		Uri:  uri,
	}
}

func (r *TextResourceContents) SetMimeType(mimeType string) *TextResourceContents {
	r.MimeType = mimeType
	return r
}

type BlobResourceContents struct {
	Blob     string `json:"blob"`
	MimeType string `json:"mimeType,omitempty"`
	Uri      string `json:"uri"`
}

func NewBlobResourceContents(blob, uri string) *BlobResourceContents {
	if blob == "" {
		panic("blob is not set")
	}
	if uri == "" {
		panic("uri is not set")
	}
	return &BlobResourceContents{
		Blob: blob,
		Uri:  uri,
	}
}
func (r *BlobResourceContents) SetMimeType(mimeType string) *BlobResourceContents {
	r.MimeType = mimeType
	return r
}

// ResourceTemplate is a template description for resources available on the server.
type ResourceTemplate struct {
	// A description of what this template is for.
	//
	// This can be used by clients to improve the LLM’s understanding of available resources. It can be thought of like a “hint” to the model.
	Description string `json:"description,omitempty"`
	MimeType    string `json:"mimeType,omitempty"`
	Name        string `json:"name"`
	Title       string `json:"title,omitempty"`
	UriTemplate string `json:"uriTemplate"`
}
