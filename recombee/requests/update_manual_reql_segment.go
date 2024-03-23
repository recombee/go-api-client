// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// UpdateManualReqlSegment Update definition of the Segment.
type UpdateManualReqlSegment struct {
	ApiRequest
	client ApiClient
}

// NewUpdateManualReqlSegment creates UpdateManualReqlSegment request.
// Update definition of the Segment.
func NewUpdateManualReqlSegment(client ApiClient, segmentationId string, segmentId string, filter string) *UpdateManualReqlSegment {

	bodyParameters := map[string]interface{}{
		"filter": filter,
	}

	queryParams := map[string]interface{}{}

	return &UpdateManualReqlSegment{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/segmentations/manual-reql/%s/segments/%s", segmentationId, segmentId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTitle sets the title parameter.
// Human-readable name of the Segment that is shown in the Recombee Admin UI.
func (r *UpdateManualReqlSegment) SetTitle(title string) *UpdateManualReqlSegment {
	r.BodyParameters["title"] = title
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *UpdateManualReqlSegment) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *UpdateManualReqlSegment) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
