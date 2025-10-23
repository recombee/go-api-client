// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItemPurchases Lists all the ever-made purchases of the given item.
type ListItemPurchases struct {
	ApiRequest
	client ApiClient
}

// NewListItemPurchases creates ListItemPurchases request.
// Lists all the ever-made purchases of the given item.
func NewListItemPurchases(client ApiClient, itemId string) *ListItemPurchases {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemPurchases{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s/purchases/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Purchase),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemPurchases) SendWithContext(ctx context.Context) ([]bindings.Purchase, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Purchase)), err
}

// Sends the request to the Recombee API
func (r *ListItemPurchases) Send() ([]bindings.Purchase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
