// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// RemoveFromSeries Removes an existing series item from the series.
type RemoveFromSeries struct {
	ApiRequest
	client ApiClient
}

// NewRemoveFromSeries creates RemoveFromSeries request.
// Removes an existing series item from the series.
func NewRemoveFromSeries(client ApiClient, seriesId string, itemType string, itemId string) *RemoveFromSeries {

	bodyParameters := map[string]interface{}{
		"itemType": itemType,
		"itemId":   itemId,
	}

	queryParams := map[string]interface{}{}

	return &RemoveFromSeries{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/series/%s/items/", seriesId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *RemoveFromSeries) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *RemoveFromSeries) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
