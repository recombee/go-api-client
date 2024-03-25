// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListSeries Gets the list of all the series currently present in the database.
type ListSeries struct {
	ApiRequest
	client ApiClient
}

// NewListSeries creates ListSeries request.
// Gets the list of all the series currently present in the database.
func NewListSeries(client ApiClient) *ListSeries {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListSeries{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/series/list/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Series),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListSeries) SendWithContext(ctx context.Context) ([]bindings.Series, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Series)), err
}

// Sends the request to the Recombee API
func (r *ListSeries) Send() ([]bindings.Series, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
