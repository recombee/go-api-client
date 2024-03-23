// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddUser Adds a new user to the database.
type AddUser struct {
	ApiRequest
	client ApiClient
}

// NewAddUser creates AddUser request.
// Adds a new user to the database.
func NewAddUser(client ApiClient, userId string) *AddUser {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &AddUser{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/users/%s", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *AddUser) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddUser) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
