package gomcp

import (
	"context"
	"net/http"
)

type HandlerResponse struct {
	Status   int
	SendBody bool
	Body     *JsonRpcResponse
}

func RequestResponse(body *JsonRpcResponse) *HandlerResponse {
	return &HandlerResponse{
		Status:   http.StatusOK,
		SendBody: true,
		Body:     body,
	}
}

func NotificationResponse() *HandlerResponse {
	return &HandlerResponse{
		Status:   http.StatusAccepted,
		SendBody: false,
		Body:     nil,
	}
}

func BadRequestResponse(body *JsonRpcResponse) *HandlerResponse {
	return &HandlerResponse{
		Status:   http.StatusBadRequest,
		SendBody: true,
		Body:     body,
	}
}

type HandleFunc func(ctx context.Context, message *JsonRpcRequest) *HandlerResponse
