// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddManualReqlSegment Adds a new Segment into a Manual ReQL Segmentation.
// The new Segment is defined by a [ReQL](https://docs.recombee.com/reql) filter that returns `true` for an item in case that this item belongs to the segment.
type AddManualReqlSegment struct {
	ApiRequest
	client ApiClient
}

// NewAddManualReqlSegment creates AddManualReqlSegment request.
// Adds a new Segment into a Manual ReQL Segmentation.
// The new Segment is defined by a [ReQL](https://docs.recombee.com/reql) filter that returns `true` for an item in case that this item belongs to the segment.
func NewAddManualReqlSegment(client ApiClient, segmentationId string, segmentId string, filter string) *AddManualReqlSegment {

	bodyParameters := map[string]interface{}{
		"filter": filter,
	}

	queryParams := map[string]interface{}{}

	return &AddManualReqlSegment{
		ApiRequest{
			HttpMethod:      http.MethodPut,
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
func (r *AddManualReqlSegment) SetTitle(title string) *AddManualReqlSegment {
	r.BodyParameters["title"] = title
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddManualReqlSegment) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddManualReqlSegment) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
