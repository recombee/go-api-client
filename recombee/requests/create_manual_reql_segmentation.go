// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// CreateManualReqlSegmentation Segment the items using multiple [ReQL](https://docs.recombee.com/reql.html) filters.
// Use the Add Manual ReQL Items Segment endpoint to create the individual segments.
type CreateManualReqlSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewCreateManualReqlSegmentation creates CreateManualReqlSegmentation request.
// Segment the items using multiple [ReQL](https://docs.recombee.com/reql.html) filters.
// Use the Add Manual ReQL Items Segment endpoint to create the individual segments.
func NewCreateManualReqlSegmentation(client ApiClient, segmentationId string, sourceType string) *CreateManualReqlSegmentation {

	bodyParameters := map[string]interface{}{
		"sourceType": sourceType,
	}

	queryParams := map[string]interface{}{}

	return &CreateManualReqlSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/segmentations/manual-reql/%s", segmentationId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTitle sets the title parameter.
// Human-readable name that is shown in the Recombee Admin UI.
func (r *CreateManualReqlSegmentation) SetTitle(title string) *CreateManualReqlSegmentation {
	r.BodyParameters["title"] = title
	return r
}

// SetDescription sets the description parameter.
// Description that is shown in the Recombee Admin UI.
func (r *CreateManualReqlSegmentation) SetDescription(description string) *CreateManualReqlSegmentation {
	r.BodyParameters["description"] = description
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *CreateManualReqlSegmentation) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *CreateManualReqlSegmentation) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
