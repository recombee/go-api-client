// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// UpdateManualReqlSegmentation Update an existing Segmentation.
type UpdateManualReqlSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewUpdateManualReqlSegmentation creates UpdateManualReqlSegmentation request.
// Update an existing Segmentation.
func NewUpdateManualReqlSegmentation(client ApiClient, segmentationId string) *UpdateManualReqlSegmentation {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &UpdateManualReqlSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodPost,
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
func (r *UpdateManualReqlSegmentation) SetTitle(title string) *UpdateManualReqlSegmentation {
	r.BodyParameters["title"] = title
	return r
}

// SetDescription sets the description parameter.
// Description that is shown in the Recombee Admin UI.
func (r *UpdateManualReqlSegmentation) SetDescription(description string) *UpdateManualReqlSegmentation {
	r.BodyParameters["description"] = description
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *UpdateManualReqlSegmentation) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *UpdateManualReqlSegmentation) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
