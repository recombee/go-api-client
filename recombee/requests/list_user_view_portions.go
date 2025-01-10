// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
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
	return *(r.ApiRequest.Target.(*[]bindings.ViewPortion)), err
}

// Sends the request to the Recombee API
func (r *ListUserViewPortions) Send() ([]bindings.ViewPortion, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
