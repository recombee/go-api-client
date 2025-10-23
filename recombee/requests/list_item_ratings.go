// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItemRatings Lists all the ratings of an item ever submitted by different users.
type ListItemRatings struct {
	ApiRequest
	client ApiClient
}

// NewListItemRatings creates ListItemRatings request.
// Lists all the ratings of an item ever submitted by different users.
func NewListItemRatings(client ApiClient, itemId string) *ListItemRatings {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItemRatings{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/%s/ratings/", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Rating),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItemRatings) SendWithContext(ctx context.Context) ([]bindings.Rating, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Rating)), err
}

// Sends the request to the Recombee API
func (r *ListItemRatings) Send() ([]bindings.Rating, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
