package requests

import (
	"context"
	"time"
)

// ApiClient defines the interface for an API client capable of sending requests.
type ApiClient interface {
	SendRequestWithContext(ctx context.Context, request *ApiRequest) error
}

// ApiRequest represents the details of a request to be sent to the API.
type ApiRequest struct {
	HttpMethod      string                 // HttpMethod is the HTTP method (e.g., "GET", "POST") for the request.
	Path            string                 // Path specifies the endpoint path of the request (including path parameters)
	QueryParameters map[string]interface{} // QueryParameters contains any query parameters to be included in the request URL.
	BodyParameters  map[string]interface{} // BodyParameters holds the parameters to be included in the request body, applicable for methods like "POST".
	DefaultTimeout  time.Duration          // DefaultTimeout specifies how long to wait for the request to complete before timing out.
	EnsureHttps     bool                   // EnsureHttps indicates whether the request should be sent over HTTPS regardless of the default client settings.
	Target          interface{}            // Target is an interface{} where the response will be decoded into.
}

func (r ApiRequest) getApiRequest() ApiRequest {
	return r
}

// Request defines an interface for types that represent API requests.
type Request interface {
	getApiRequest() ApiRequest
}
