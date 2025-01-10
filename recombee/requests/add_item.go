// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddItem Adds new item of the given `itemId` to the items catalog.
// All the item properties for the newly created items are set to null.
type AddItem struct {
	ApiRequest
	client ApiClient
}

// NewAddItem creates AddItem request.
// Adds new item of the given `itemId` to the items catalog.
// All the item properties for the newly created items are set to null.
func NewAddItem(client ApiClient, itemId string) *AddItem {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &AddItem{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/items/%s", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *AddItem) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddItem) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
