package gomcp

import (
	"encoding/json"
	"io"
)

type JsonRpcRequest struct {
	Jsonrpc string          `json:"jsonrpc"`
	Method  string          `json:"method"`
	Params  json.RawMessage `json:"params"`
	Id      any             `json:"id"`
}

func ReadJsonRpcRequest(r io.Reader) (*JsonRpcRequest, error) {
	var request JsonRpcRequest
	err := json.NewDecoder(r).Decode(&request)
	return &request, err
}

type JsonRpcResponse struct {
	Jsonrpc string        `json:"jsonrpc"`
	Result  any           `json:"result,omitempty"`
	Error   *JsonRpcError `json:"error,omitempty"`
	Id      any           `json:"id"`
}

func NewResultJsonRpcResponse(id, result any) *JsonRpcResponse {
	return &JsonRpcResponse{
		Jsonrpc: "2.0",
		Result:  result,
		Id:      id,
	}
}

func NewErrorJsonRpcResponse(id any, error *JsonRpcError) *JsonRpcResponse {
	return &JsonRpcResponse{
		Jsonrpc: "2.0",
		Error:   error,
		Id:      id,
	}
}

func (r *JsonRpcResponse) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(r)
}

type JsonRpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}
