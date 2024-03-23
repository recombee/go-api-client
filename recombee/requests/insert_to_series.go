// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// InsertToSeries Inserts an existing item/series into a series of the given seriesId at a position determined by time.
type InsertToSeries struct {
	ApiRequest
	client ApiClient
}

// NewInsertToSeries creates InsertToSeries request.
// Inserts an existing item/series into a series of the given seriesId at a position determined by time.
func NewInsertToSeries(client ApiClient, seriesId string, itemType string, itemId string, time float64) *InsertToSeries {

	bodyParameters := map[string]interface{}{
		"itemType": itemType,
		"itemId":   itemId,
		"time":     time,
	}

	queryParams := map[string]interface{}{}

	return &InsertToSeries{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/series/%s/items/", seriesId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Indicates that any non-existing entity specified within the request should be created (as if corresponding PUT requests were invoked). This concerns both the `seriesId` and the `itemId`. If `cascadeCreate` is set to true, the behavior also depends on the `itemType`. Either item or series may be created if not present in the database.
func (r *InsertToSeries) SetCascadeCreate(cascadeCreate bool) *InsertToSeries {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *InsertToSeries) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *InsertToSeries) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
