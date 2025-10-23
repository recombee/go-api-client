// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListUserProperties Gets the list of all the user properties in your database.
type ListUserProperties struct {
	ApiRequest
	client ApiClient
}

// NewListUserProperties creates ListUserProperties request.
// Gets the list of all the user properties in your database.
func NewListUserProperties(client ApiClient) *ListUserProperties {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserProperties{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/properties/list/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.PropertyInfo),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserProperties) SendWithContext(ctx context.Context) ([]bindings.PropertyInfo, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.PropertyInfo)), err
}

// Sends the request to the Recombee API
func (r *ListUserProperties) Send() ([]bindings.PropertyInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
