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
