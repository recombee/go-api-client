// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// GetUserPropertyInfo Gets information about specified user property.
type GetUserPropertyInfo struct {
	ApiRequest
	client ApiClient
}

// NewGetUserPropertyInfo creates GetUserPropertyInfo request.
// Gets information about specified user property.
func NewGetUserPropertyInfo(client ApiClient, propertyName string) *GetUserPropertyInfo {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &GetUserPropertyInfo{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/properties/%s", propertyName),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(bindings.PropertyInfo),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *GetUserPropertyInfo) SendWithContext(ctx context.Context) (bindings.PropertyInfo, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.PropertyInfo{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.PropertyInfo)), err
}

// Sends the request to the Recombee API
func (r *GetUserPropertyInfo) Send() (bindings.PropertyInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
