// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// GetItemValues Gets all the current property values of the given item.
type GetItemValues struct {
	ApiRequest
	client ApiClient
}

// NewGetItemValues creates GetItemValues request.
// Gets all the current property values of the given item.
func NewGetItemValues(client ApiClient, itemId string) *GetItemValues {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &GetItemValues{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(map[string]interface{}),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *GetItemValues) SendWithContext(ctx context.Context) (map[string]interface{}, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return *(r.ApiRequest.Target.(*map[string]interface{})), err
}

// Sends the request to the Recombee API
func (r *GetItemValues) Send() (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
