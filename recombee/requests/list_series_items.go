// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListSeriesItems Lists all the items present in the given series, sorted according to their time index values.
type ListSeriesItems struct {
	ApiRequest
	client ApiClient
}

// NewListSeriesItems creates ListSeriesItems request.
// Lists all the items present in the given series, sorted according to their time index values.
func NewListSeriesItems(client ApiClient, seriesId string) *ListSeriesItems {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListSeriesItems{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/series/%s/items/", seriesId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.SeriesItem),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListSeriesItems) SendWithContext(ctx context.Context) ([]bindings.SeriesItem, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.SeriesItem)), err
}

// Sends the request to the Recombee API
func (r *ListSeriesItems) Send() ([]bindings.SeriesItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
