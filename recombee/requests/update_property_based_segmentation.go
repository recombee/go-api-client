// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// UpdatePropertyBasedSegmentation Updates a Property Based Segmentation
type UpdatePropertyBasedSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewUpdatePropertyBasedSegmentation creates UpdatePropertyBasedSegmentation request.
// Updates a Property Based Segmentation
func NewUpdatePropertyBasedSegmentation(client ApiClient, segmentationId string) *UpdatePropertyBasedSegmentation {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &UpdatePropertyBasedSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/segmentations/property-based/%s", segmentationId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetPropertyName sets the propertyName parameter.
// Name of the property on which the Segmentation should be based
func (r *UpdatePropertyBasedSegmentation) SetPropertyName(propertyName string) *UpdatePropertyBasedSegmentation {
	r.BodyParameters["propertyName"] = propertyName
	return r
}

// SetTitle sets the title parameter.
// Human-readable name that is shown in the Recombee Admin UI.
func (r *UpdatePropertyBasedSegmentation) SetTitle(title string) *UpdatePropertyBasedSegmentation {
	r.BodyParameters["title"] = title
	return r
}

// SetDescription sets the description parameter.
// Description that is shown in the Recombee Admin UI.
func (r *UpdatePropertyBasedSegmentation) SetDescription(description string) *UpdatePropertyBasedSegmentation {
	r.BodyParameters["description"] = description
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *UpdatePropertyBasedSegmentation) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *UpdatePropertyBasedSegmentation) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
