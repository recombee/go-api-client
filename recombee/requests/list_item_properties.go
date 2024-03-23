// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItemProperties Gets the list of all the item properties in your database.
type ListItemProperties struct {
	ApiRequest
	client ApiClient
}

// NewListItemProperties creates ListItemProperties request.
// Gets the list of all the item properties in your database.
func NewListItemProperties(client ApiClient) *ListItemProperties {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemProperties{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/properties/list/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.PropertyInfo),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemProperties) SendWithContext(ctx context.Context) ([]bindings.PropertyInfo, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.PropertyInfo)), err
}

// Sends the request to the Recombee API
func (r *ListItemProperties) Send() ([]bindings.PropertyInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
