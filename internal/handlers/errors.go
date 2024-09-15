package handlers

import (
	"fmt"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func NewErrorResponse(status int, error, message string, details interface{}) *Response {
	return &Response{
		Status:  status,
		Error:   error,
		Message: error,
		Details: details,
	}
}

func SendError(status int, error, message string, details interface{}) error {
	errResponse := NewErrorResponse(status, error, message, details)
	return fmt.Errorf(
		"%d\n, %s\n, %s\n, %v\n",
		errResponse.Status,
		errResponse.Error,
		errResponse.Message,
		errResponse.Details)
}

func OK() *Response {
	return &Response{
		Status: http.StatusOK,
	}
}
