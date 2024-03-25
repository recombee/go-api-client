// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
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
			Path:            fmt.Sprintf("/segmentations/list/"),
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
	return *(r.ApiRequest.Target.(*bindings.ListSegmentationsResponse)), err
}

// Sends the request to the Recombee API
func (r *ListSegmentations) Send() (bindings.ListSegmentationsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
