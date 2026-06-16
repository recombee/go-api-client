// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name

	"github.com/recombee/go-api-client/v6/recombee/bindings"
)

// ListItemViewPortions Lists all the view portions of an item ever submitted by different users.
type ListItemViewPortions struct {
	ApiRequest
	client ApiClient
}

// NewListItemViewPortions creates ListItemViewPortions request.
// Lists all the view portions of an item ever submitted by different users.
func NewListItemViewPortions(client ApiClient, itemId string) *ListItemViewPortions {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemViewPortions{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s/viewportions/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.ViewPortion),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemViewPortions) SendWithContext(ctx context.Context) ([]bindings.ViewPortion, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.Target.(*[]bindings.ViewPortion)), err
}

// Sends the request to the Recombee API
func (r *ListItemViewPortions) Send() ([]bindings.ViewPortion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
