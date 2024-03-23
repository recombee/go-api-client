// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddUserProperty Adding a user property is somewhat equivalent to adding a column to the table of users. The users may be characterized by various properties of different types.
type AddUserProperty struct {
	ApiRequest
	client ApiClient
}

// NewAddUserProperty creates AddUserProperty request.
// Adding a user property is somewhat equivalent to adding a column to the table of users. The users may be characterized by various properties of different types.
func NewAddUserProperty(client ApiClient, propertyName string, propertyType string) *AddUserProperty {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"type": propertyType,
	}

	return &AddUserProperty{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/users/properties/%s", propertyName),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *AddUserProperty) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddUserProperty) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
