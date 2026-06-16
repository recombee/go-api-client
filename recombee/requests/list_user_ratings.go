// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name

	"github.com/recombee/go-api-client/v6/recombee/bindings"
)

// ListUserRatings Lists all the ratings ever submitted by the given user.
type ListUserRatings struct {
	ApiRequest
	client ApiClient
}

// NewListUserRatings creates ListUserRatings request.
// Lists all the ratings ever submitted by the given user.
func NewListUserRatings(client ApiClient, userId string) *ListUserRatings {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserRatings{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s/ratings/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Rating),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserRatings) SendWithContext(ctx context.Context) ([]bindings.Rating, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.Target.(*[]bindings.Rating)), err
}

// Sends the request to the Recombee API
func (r *ListUserRatings) Send() ([]bindings.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
