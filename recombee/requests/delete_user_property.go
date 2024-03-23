// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteUserProperty Deleting a user property is roughly equivalent to removing a column from the table of users.
type DeleteUserProperty struct {
	ApiRequest
	client ApiClient
}

// NewDeleteUserProperty creates DeleteUserProperty request.
// Deleting a user property is roughly equivalent to removing a column from the table of users.
func NewDeleteUserProperty(client ApiClient, propertyName string) *DeleteUserProperty {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteUserProperty{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
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
func (r *DeleteUserProperty) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteUserProperty) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
