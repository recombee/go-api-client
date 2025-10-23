// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListUserCartAdditions Lists all the cart additions ever made by the given user.
type ListUserCartAdditions struct {
	ApiRequest
	client ApiClient
}

// NewListUserCartAdditions creates ListUserCartAdditions request.
// Lists all the cart additions ever made by the given user.
func NewListUserCartAdditions(client ApiClient, userId string) *ListUserCartAdditions {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserCartAdditions{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s/cartadditions/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.CartAddition),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserCartAdditions) SendWithContext(ctx context.Context) ([]bindings.CartAddition, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.CartAddition)), err
}

// Sends the request to the Recombee API
func (r *ListUserCartAdditions) Send() ([]bindings.CartAddition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
