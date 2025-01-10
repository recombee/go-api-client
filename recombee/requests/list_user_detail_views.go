// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListUserDetailViews Lists all the detail views of different items ever made by the given user.
type ListUserDetailViews struct {
	ApiRequest
	client ApiClient
}

// NewListUserDetailViews creates ListUserDetailViews request.
// Lists all the detail views of different items ever made by the given user.
func NewListUserDetailViews(client ApiClient, userId string) *ListUserDetailViews {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserDetailViews{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s/detailviews/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.DetailView),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserDetailViews) SendWithContext(ctx context.Context) ([]bindings.DetailView, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.DetailView)), err
}

// Sends the request to the Recombee API
func (r *ListUserDetailViews) Send() ([]bindings.DetailView, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
