package requests

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// SetItemValues Sets/updates (some) property values of the given item. The properties (columns) must be previously created by [Add item property](https://docs.recombee.com/api.html#add-item-property).
type SetItemValues struct {
	ApiRequest
	client ApiClient
}

// NewSetItemValues creates SetItemValues request.
//
// Sets/updates (some) property values of the given item. The properties (columns) must be previously created by [Add item property](https://docs.recombee.com/api.html#add-item-property).
func NewSetItemValues(client ApiClient, itemId string, values map[string]interface{}) *SetItemValues {

	queryParams := map[string]interface{}{}

	return &SetItemValues{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/items/%s", itemId),
			BodyParameters:  values,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * time.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the item should be created if not present in the database.
func (r *SetItemValues) SetCascadeCreate(cascadeCreate bool) *SetItemValues {
	r.BodyParameters["!cascadeCreate"] = cascadeCreate
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *SetItemValues) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *SetItemValues) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
