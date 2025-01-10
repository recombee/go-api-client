package errors

import (
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/requests"
)

// ResponseError is returned when request did not succeed (did not return 200 or 201)
type ResponseError struct {
	// HTTP status code returned by the API
	StatusCode int

	// Error message returned by the API
	Message string

	// ApiRequest which failed
	ApiRequest *requests.ApiRequest
}

// NewResponseError creates and returns a new instance of ResponseError.
func NewResponseError(statusCode int, message string, request *requests.ApiRequest) *ResponseError {
	return &ResponseError{
		StatusCode: statusCode,
		Message:    message,
		ApiRequest: request,
	}
}

// Error implements the error interface for ResponseError.
func (e *ResponseError) Error() string {
	return fmt.Sprintf("Status %d: %s", e.StatusCode, e.Message)
}
