// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
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

	bodyParameters := map[string]interface{}{
		"type": propertyType,
	}

	queryParams := map[string]interface{}{}

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

// SetRole sets the role parameter.
// [Role](https://docs.recombee.com/api/property_roles_metadata#roles) to assign to the property.
func (r *AddItemProperty) SetRole(role bindings.PropertyRole) *AddItemProperty {
	r.BodyParameters["role"] = role
	return r
}

// SetMetadata sets the metadata parameter.
// List of [metadata](https://docs.recombee.com/api/property_roles_metadata#metadata) entries to assign to the property.
func (r *AddItemProperty) SetMetadata(metadata []bindings.PropertyMetadata) *AddItemProperty {
	r.BodyParameters["metadata"] = metadata
	return r
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
