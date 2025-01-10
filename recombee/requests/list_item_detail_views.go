// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItemDetailViews Lists all the detail views of the given item ever made by different users.
type ListItemDetailViews struct {
	ApiRequest
	client ApiClient
}

// NewListItemDetailViews creates ListItemDetailViews request.
// Lists all the detail views of the given item ever made by different users.
func NewListItemDetailViews(client ApiClient, itemId string) *ListItemDetailViews {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemDetailViews{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s/detailviews/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.DetailView),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemDetailViews) SendWithContext(ctx context.Context) ([]bindings.DetailView, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.DetailView)), err
}

// Sends the request to the Recombee API
func (r *ListItemDetailViews) Send() ([]bindings.DetailView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
