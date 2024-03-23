// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddItemProperty Adding an item property is somewhat equivalent to adding a column to the table of items. The items may be characterized by various properties of different types.
type AddItemProperty struct {
	ApiRequest
	client ApiClient
}

// NewAddItemProperty creates AddItemProperty request.
// Adding an item property is somewhat equivalent to adding a column to the table of items. The items may be characterized by various properties of different types.
func NewAddItemProperty(client ApiClient, propertyName string, propertyType string) *AddItemProperty {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"type": propertyType,
	}

	return &AddItemProperty{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/items/properties/%s", propertyName),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *AddItemProperty) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddItemProperty) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
