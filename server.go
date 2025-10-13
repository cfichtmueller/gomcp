package gomcp

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/cfichtmueller/gomcp/protocol"
)

type Server struct {
	tools     []*Tool
	resources []*Resource
	handlers  map[string]HandleFunc
}

func NewServer() *Server {
	s := &Server{
		tools:     make([]*Tool, 0),
		resources: make([]*Resource, 0),
		handlers:  make(map[string]HandleFunc),
	}

	s.handlers["initialize"] = s.handleInitialize
	s.handlers["logging/setLevel"] = s.handleLoggingSetLevel
	s.handlers["notifications/initialized"] = s.handleInitializedNotification
	s.handlers["resources/list"] = s.handleListResources
	s.handlers["resources/read"] = s.handleReadResource
	s.handlers["tools/call"] = s.handleCallTool
	s.handlers["tools/list"] = s.handleListTools
	return s
}

func (s *Server) AddResource(resource *Resource) {
	if resource.Name == "" {
		panic("name is not set")
	}
	if resource.Uri == "" {
		panic("uri is not set")
	}
	s.resources = append(s.resources, resource)
}

func (s *Server) AddTool(tool *Tool) {
	if tool.Name == "" {
		panic("name is not set")
	}
	if tool.InputSchema == nil {
		panic("input schema is not set")
	}
	s.tools = append(s.tools, tool)
}

func (s *Server) handle(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	handler, ok := s.handlers[message.Method]
	if !ok {
		return BadRequestResponse(&JsonRpcResponse{
			Jsonrpc: "2.0",
			Error: &JsonRpcError{
				Code:    -32000,
				Message: "Unsupported method",
			},
		})
	}

	return handler(ctx, message)
}

func (s *Server) handleInitialize(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	var params protocol.InitializeParams
	if r := s.mustParseParams(message, &params); r != nil {
		return r
	}

	slog.Info("Initialization started", "params", params)
	caps := protocol.NewServerCapabilities()
	caps.Tools = protocol.NewCapability().SetListChanged(true)
	if len(s.resources) > 0 {
		caps.Resources = protocol.NewCapability().SetListChanged(true)
	}
	return RequestResponse(NewResultJsonRpcResponse(message.Id, protocol.InitializeResult{
		ProtocolVersion: params.ProtocolVersion,
		Capabilities:    caps,
		ServerInfo: &protocol.ServerInfo{
			Name:    "gomcp",
			Title:   "gomcp",
			Version: Release,
		},
		Instructions: "Hello, world!",
	}))
}

func (s *Server) handleInitializedNotification(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	return NotificationResponse()
}

func (s *Server) handleLoggingSetLevel(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	var params protocol.LoggingSetLevelParams
	if r := s.mustParseParams(message, &params); r != nil {
		return r
	}

	slog.Info("Logging set level", "params", params)
	return NotificationResponse()
}

func (s *Server) handleListResources(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	var params protocol.ListResourcesRequest
	if r := s.mustParseParams(message, &params); r != nil {
		return r
	}
	res := protocol.NewListResourcesResult()
	for _, resource := range s.resources {
		res.AddResource(&protocol.Resource{
			Name: resource.Name,
			Uri:  resource.Uri,
		})
	}
	return RequestResponse(NewResultJsonRpcResponse(message.Id, res))
}

func (s *Server) handleReadResource(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	var params protocol.ReadResourceParams
	if r := s.mustParseParams(message, &params); r != nil {
		return r
	}
	for _, resource := range s.resources {
		if resource.Uri == params.Uri {
			r := resource.Handler(ctx)
			return RequestResponse(NewResultJsonRpcResponse(message.Id, r))
		}
	}
	return BadRequestResponse(NewErrorJsonRpcResponse(message.Id, &JsonRpcError{
		Code:    -32000,
		Message: "Resource not found",
	}))
}

func (s *Server) handleCallTool(ctx context.Context, message *JsonRpcRequest) *HandlerResponse {
	var params protocol.CallToolsParams
	if r := s.mustParseParams(message, &params); r != nil {
		return r
	}
	var tool *Tool
	for _, t := range s.tools {
		if t.Name == params.Name {
			tool = t
			break
		}

	}
	if tool == nil {
		return BadRequestResponse(NewErrorJsonRpcResponse(message.Id, &JsonRpcError{
			Code:    -32000,
			Message: "Tool not found",
		}))
	}
	args := NewToolArguments(params.Arguments)
	return RequestResponse(NewResultJsonRpcResponse(message.Id, tool.Call(ctx, args)))

}

func (s *Server) handleListTools(ctx context.Context, request *JsonRpcRequest) *HandlerResponse {
	res := protocol.NewListToolsResult()
	for _, tool := range s.tools {
		res.AddTool(&protocol.Tool{
			Name:         tool.Name,
			Title:        tool.Title,
			Description:  tool.Description,
			InputSchema:  tool.InputSchema,
			OutputSchema: tool.OutputSchema,
		})
	}
	return RequestResponse(NewResultJsonRpcResponse(request.Id, res))
}

func (s *Server) mustParseParams(message *JsonRpcRequest, params any) *HandlerResponse {
	if err := json.Unmarshal(message.Params, &params); err != nil {
		return BadRequestResponse(NewErrorJsonRpcResponse(message.Id, &JsonRpcError{
			Code:    -32000,
			Message: "Invalid params",
		}))
	}

	return nil
}
