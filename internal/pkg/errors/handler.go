package errors

import "github.com/labstack/echo/v4"

type ErrorResponse struct {
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func NewErrorResponse(status int, error, message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Status:  status,
		Error:   error,
		Message: error,
		Details: details,
	}
}

func SendError(c echo.Context, status int, error, message string, details interface{}) error {
	errResponse := NewErrorResponse(status, error, message, details)
	return c.JSON(status, errResponse)
}
