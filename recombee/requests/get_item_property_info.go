// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name

	"github.com/recombee/go-api-client/v6/recombee/bindings"
)

// GetItemPropertyInfo Gets information about specified item property.
type GetItemPropertyInfo struct {
	ApiRequest
	client ApiClient
}

// NewGetItemPropertyInfo creates GetItemPropertyInfo request.
// Gets information about specified item property.
func NewGetItemPropertyInfo(client ApiClient, propertyName string) *GetItemPropertyInfo {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &GetItemPropertyInfo{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/properties/%s", propertyName),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(bindings.PropertyInfo),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *GetItemPropertyInfo) SendWithContext(ctx context.Context) (bindings.PropertyInfo, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.PropertyInfo{}, err
	}
	return *(r.Target.(*bindings.PropertyInfo)), err
}

// Sends the request to the Recombee API
func (r *GetItemPropertyInfo) Send() (bindings.PropertyInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
