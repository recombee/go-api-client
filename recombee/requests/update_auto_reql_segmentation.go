// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// UpdateAutoReqlSegmentation Update an existing Segmentation.
type UpdateAutoReqlSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewUpdateAutoReqlSegmentation creates UpdateAutoReqlSegmentation request.
// Update an existing Segmentation.
func NewUpdateAutoReqlSegmentation(client ApiClient, segmentationId string) *UpdateAutoReqlSegmentation {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &UpdateAutoReqlSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/segmentations/auto-reql/%s", segmentationId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetExpression sets the expression parameter.
// ReQL expression that returns for each item a set with IDs of segments to which the item belongs
func (r *UpdateAutoReqlSegmentation) SetExpression(expression string) *UpdateAutoReqlSegmentation {
	r.BodyParameters["expression"] = expression
	return r
}

// SetTitle sets the title parameter.
// Human-readable name that is shown in the Recombee Admin UI.
func (r *UpdateAutoReqlSegmentation) SetTitle(title string) *UpdateAutoReqlSegmentation {
	r.BodyParameters["title"] = title
	return r
}

// SetDescription sets the description parameter.
// Description that is shown in the Recombee Admin UI.
func (r *UpdateAutoReqlSegmentation) SetDescription(description string) *UpdateAutoReqlSegmentation {
	r.BodyParameters["description"] = description
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *UpdateAutoReqlSegmentation) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *UpdateAutoReqlSegmentation) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
