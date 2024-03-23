// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// CreatePropertyBasedSegmentation Creates a Segmentation that splits the items into segments based on values of a particular item property.
// A segment is created for each unique value of the property.
// In case of `set` properties, a segment is created for each value in the set. Item belongs to all these segments.
type CreatePropertyBasedSegmentation struct {
	ApiRequest
	client ApiClient
}

// NewCreatePropertyBasedSegmentation creates CreatePropertyBasedSegmentation request.
// Creates a Segmentation that splits the items into segments based on values of a particular item property.
// A segment is created for each unique value of the property.
// In case of `set` properties, a segment is created for each value in the set. Item belongs to all these segments.
func NewCreatePropertyBasedSegmentation(client ApiClient, segmentationId string, sourceType string, propertyName string) *CreatePropertyBasedSegmentation {

	bodyParameters := map[string]interface{}{
		"sourceType":   sourceType,
		"propertyName": propertyName,
	}

	queryParams := map[string]interface{}{}

	return &CreatePropertyBasedSegmentation{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/segmentations/property-based/%s", segmentationId),
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
func (r *CreatePropertyBasedSegmentation) SetTitle(title string) *CreatePropertyBasedSegmentation {
	r.BodyParameters["title"] = title
	return r
}

// SetDescription sets the description parameter.
// Description that is shown in the Recombee Admin UI.
func (r *CreatePropertyBasedSegmentation) SetDescription(description string) *CreatePropertyBasedSegmentation {
	r.BodyParameters["description"] = description
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *CreatePropertyBasedSegmentation) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *CreatePropertyBasedSegmentation) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
