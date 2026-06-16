// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name

	"github.com/recombee/go-api-client/v6/recombee/bindings"
)

// ListUserViewPortions Lists all the view portions ever submitted by the given user.
type ListUserViewPortions struct {
	ApiRequest
	client ApiClient
}

// NewListUserViewPortions creates ListUserViewPortions request.
// Lists all the view portions ever submitted by the given user.
func NewListUserViewPortions(client ApiClient, userId string) *ListUserViewPortions {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserViewPortions{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s/viewportions/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.ViewPortion),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserViewPortions) SendWithContext(ctx context.Context) ([]bindings.ViewPortion, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.Target.(*[]bindings.ViewPortion)), err
}

// Sends the request to the Recombee API
func (r *ListUserViewPortions) Send() ([]bindings.ViewPortion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
