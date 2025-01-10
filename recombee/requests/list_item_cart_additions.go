// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItemCartAdditions Lists all the ever-made cart additions of the given item.
type ListItemCartAdditions struct {
	ApiRequest
	client ApiClient
}

// NewListItemCartAdditions creates ListItemCartAdditions request.
// Lists all the ever-made cart additions of the given item.
func NewListItemCartAdditions(client ApiClient, itemId string) *ListItemCartAdditions {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemCartAdditions{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s/cartadditions/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.CartAddition),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemCartAdditions) SendWithContext(ctx context.Context) ([]bindings.CartAddition, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.CartAddition)), err
}

// Sends the request to the Recombee API
func (r *ListItemCartAdditions) Send() ([]bindings.CartAddition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
