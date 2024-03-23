package bindings

// BatchResponseError provides additional information about the failed request
type BatchResponseError struct {
	// ErrorMessage returned by the API
	ErrorMessage string `json:"error,omitempty"`
}

// BatchResponse is response for a single request sent in a Batch
type BatchResponse struct {
	// HTTP status code of the sub-request
	StatusCode int

	// Result in case of a successful request
	Result interface{}

	// Error in case the request did not succeed
	Error *BatchResponseError
}
