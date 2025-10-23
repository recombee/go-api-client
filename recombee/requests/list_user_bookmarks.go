// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListUserBookmarks Lists all the bookmarks ever made by the given user.
type ListUserBookmarks struct {
	ApiRequest
	client ApiClient
}

// NewListUserBookmarks creates ListUserBookmarks request.
// Lists all the bookmarks ever made by the given user.
func NewListUserBookmarks(client ApiClient, userId string) *ListUserBookmarks {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserBookmarks{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s/bookmarks/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Bookmark),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserBookmarks) SendWithContext(ctx context.Context) ([]bindings.Bookmark, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Bookmark)), err
}

// Sends the request to the Recombee API
func (r *ListUserBookmarks) Send() ([]bindings.Bookmark, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
