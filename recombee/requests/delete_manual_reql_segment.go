// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteManualReqlSegment Delete a Segment from a Manual ReQL Segmentation.
type DeleteManualReqlSegment struct {
	ApiRequest
	client ApiClient
}

// NewDeleteManualReqlSegment creates DeleteManualReqlSegment request.
// Delete a Segment from a Manual ReQL Segmentation.
func NewDeleteManualReqlSegment(client ApiClient, segmentationId string, segmentId string) *DeleteManualReqlSegment {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteManualReqlSegment{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteManualReqlSegment) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteManualReqlSegment) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
