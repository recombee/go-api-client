// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListUserPurchases Lists all the purchases ever made by the given user.
type ListUserPurchases struct {
	ApiRequest
	client ApiClient
}

// NewListUserPurchases creates ListUserPurchases request.
// Lists all the purchases ever made by the given user.
func NewListUserPurchases(client ApiClient, userId string) *ListUserPurchases {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUserPurchases{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/%s/purchases/", userId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Purchase),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUserPurchases) SendWithContext(ctx context.Context) ([]bindings.Purchase, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Purchase)), err
}

// Sends the request to the Recombee API
func (r *ListUserPurchases) Send() ([]bindings.Purchase, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
