// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ResetDatabase Completely erases all your data, including items, item properties, series, user database, purchases, ratings, detail views, and bookmarks. Make sure the request is never executed in the production environment! Resetting your database is irreversible.
type ResetDatabase struct {
	ApiRequest
	client ApiClient
}

// NewResetDatabase creates ResetDatabase request.
// Completely erases all your data, including items, item properties, series, user database, purchases, ratings, detail views, and bookmarks. Make sure the request is never executed in the production environment! Resetting your database is irreversible.
func NewResetDatabase(client ApiClient) *ResetDatabase {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ResetDatabase{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ResetDatabase) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *ResetDatabase) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
