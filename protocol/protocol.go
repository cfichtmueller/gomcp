package protocol

type InitializeParams struct {
	ProtocolVersion string              `json:"protocolVersion"`
	Capabilities    *ClientCapabilities `json:"capabilities"`
	ClientInfo      *ClientInfo         `json:"clientInfo"`
}

type ClientCapabilities struct {
	Methods map[string]any `json:"methods"`
}

type ClientInfo struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Version string `json:"version"`
}

type InitializeResult struct {
	ProtocolVersion string              `json:"protocolVersion"`
	Capabilities    *ServerCapabilities `json:"capabilities"`
	ServerInfo      *ServerInfo         `json:"serverInfo"`
	Instructions    string              `json:"instructions,omitempty"`
}

type ServerCapabilities struct {
	Logging   *Capability `json:"logging,omitempty"`
	Prompts   *Capability `json:"prompts,omitempty"`
	Resources *Capability `json:"resources,omitempty"`
	Tools     *Capability `json:"tools,omitempty"`
}

func NewServerCapabilities() *ServerCapabilities {
	return &ServerCapabilities{}
}

type Capability struct {
	Subscribed  *bool `json:"subscribed,omitempty"`
	ListChanged *bool `json:"listChanged,omitempty"`
}

func NewCapability() *Capability {
	return &Capability{
		Subscribed:  nil,
		ListChanged: nil,
	}
}

func (c *Capability) SetSubscribed(subscribed bool) {
	c.Subscribed = &subscribed
}

func (c *Capability) SetListChanged(listChanged bool) *Capability {
	c.ListChanged = &listChanged
	return c
}

type ServerInfo struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Version string `json:"version"`
}

type LoggingSetLevelParams struct {
	Level string `json:"level"`
}
