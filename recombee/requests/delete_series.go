// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteSeries Deletes the series of the given `seriesId` from the database.
// Deleting a series will only delete assignment of items to it, not the items themselves!
type DeleteSeries struct {
	ApiRequest
	client ApiClient
}

// NewDeleteSeries creates DeleteSeries request.
// Deletes the series of the given `seriesId` from the database.
// Deleting a series will only delete assignment of items to it, not the items themselves!
func NewDeleteSeries(client ApiClient, seriesId string) *DeleteSeries {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteSeries{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/series/%s", seriesId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteSeries) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteSeries) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
