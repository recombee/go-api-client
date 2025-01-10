// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItemBookmarks Lists all the ever-made bookmarks of the given item.
type ListItemBookmarks struct {
	ApiRequest
	client ApiClient
}

// NewListItemBookmarks creates ListItemBookmarks request.
// Lists all the ever-made bookmarks of the given item.
func NewListItemBookmarks(client ApiClient, itemId string) *ListItemBookmarks {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemBookmarks{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s/bookmarks/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Bookmark),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemBookmarks) SendWithContext(ctx context.Context) ([]bindings.Bookmark, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Bookmark)), err
}

// Sends the request to the Recombee API
func (r *ListItemBookmarks) Send() ([]bindings.Bookmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
