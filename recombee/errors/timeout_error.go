package errors

import (
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/requests"
	"time"
)

// TimeoutError represents an error that occurs when an API request is not completed within the expected time frame.
// This timeout be caused by network connection issues, or by issues of the API itself.
type TimeoutError struct {

	// ElapsedTime is the duration from when the request was initiated to when the timeout occurred.
	ElapsedTime time.Duration

	// ApiRequest is the pointer to the original request that triggered the timeout.
	ApiRequest *requests.ApiRequest

	// OriginalError holds the original error returned by the Go HTTP client when the timeout occurred.
	OriginalError error
}

// Error implements the error interface for TimeoutError.
func (e *TimeoutError) Error() string {
	return fmt.Sprintf("Timeout error after %v ms: %v", e.ElapsedTime.Milliseconds(), e.OriginalError.Error())
}

// Unwrap provides the original error, allowing errors.Is and errors.As to work.
func (e *TimeoutError) Unwrap() error {
	return e.OriginalError
}

// NewTimeoutError creates and returns a new instance of TimeoutError.
func NewTimeoutError(elapsedTime time.Duration, request *requests.ApiRequest, err error) *TimeoutError {
	return &TimeoutError{
		ElapsedTime:   elapsedTime,
		ApiRequest:    request,
		OriginalError: err,
	}
}
