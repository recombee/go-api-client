// This file is auto-generated

package requests

import (
	"context"
	"net/http"
	timepkg "time" // avoid collision with param name

	"github.com/recombee/go-api-client/v6/recombee/bindings"
)

// ListSegmentations Return all existing items Segmentations.
type ListSegmentations struct {
	ApiRequest
	client ApiClient
}

// NewListSegmentations creates ListSegmentations request.
// Return all existing items Segmentations.
func NewListSegmentations(client ApiClient, sourceType string) *ListSegmentations {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"sourceType": sourceType,
	}

	return &ListSegmentations{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            "/segmentations/list/",
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(bindings.ListSegmentationsResponse),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListSegmentations) SendWithContext(ctx context.Context) (bindings.ListSegmentationsResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.ListSegmentationsResponse{}, err
	}
	return *(r.Target.(*bindings.ListSegmentationsResponse)), err
}

// Sends the request to the Recombee API
func (r *ListSegmentations) Send() (bindings.ListSegmentationsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
