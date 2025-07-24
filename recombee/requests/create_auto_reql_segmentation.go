// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// CreateAutoReqlSegmentation Segment the items using a [ReQL](https://docs.recombee.com/reql) expression.
// For each item, the expression should return a set that contains IDs of segments to which the item belongs to.
type CreateAutoReqlSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewCreateAutoReqlSegmentation creates CreateAutoReqlSegmentation request.
// Segment the items using a [ReQL](https://docs.recombee.com/reql) expression.
// For each item, the expression should return a set that contains IDs of segments to which the item belongs to.
func NewCreateAutoReqlSegmentation(client ApiClient, segmentationId string, sourceType string, expression string) *CreateAutoReqlSegmentation {

	bodyParameters := map[string]interface{}{
		"sourceType": sourceType,
		"expression": expression,
	}

	queryParams := map[string]interface{}{}

	return &CreateAutoReqlSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/segmentations/auto-reql/%s", segmentationId),
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
func (r *CreateAutoReqlSegmentation) SetTitle(title string) *CreateAutoReqlSegmentation {
	r.BodyParameters["title"] = title
	return r
}

// SetDescription sets the description parameter.
// Description that is shown in the Recombee Admin UI.
func (r *CreateAutoReqlSegmentation) SetDescription(description string) *CreateAutoReqlSegmentation {
	r.BodyParameters["description"] = description
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *CreateAutoReqlSegmentation) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *CreateAutoReqlSegmentation) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
