// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteUser Deletes a user of the given *userId* from the database.
// If there are any purchases, ratings, bookmarks, cart additions or detail views made by the user present in the database, they will be deleted in cascade as well.
type DeleteUser struct {
	ApiRequest
	client ApiClient
}

// NewDeleteUser creates DeleteUser request.
// Deletes a user of the given *userId* from the database.
// If there are any purchases, ratings, bookmarks, cart additions or detail views made by the user present in the database, they will be deleted in cascade as well.
func NewDeleteUser(client ApiClient, userId string) *DeleteUser {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteUser{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
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
func (r *DeleteUser) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteUser) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
