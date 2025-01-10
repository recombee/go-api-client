// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// GetUserValues Gets all the current property values of the given user.
type GetUserValues struct {
	ApiRequest
	client ApiClient
}

// NewGetUserValues creates GetUserValues request.
// Gets all the current property values of the given user.
func NewGetUserValues(client ApiClient, userId string) *GetUserValues {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &GetUserValues{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(map[string]interface{}),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *GetUserValues) SendWithContext(ctx context.Context) (map[string]interface{}, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return *(r.ApiRequest.Target.(*map[string]interface{})), err
}

// Sends the request to the Recombee API
func (r *GetUserValues) Send() (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
