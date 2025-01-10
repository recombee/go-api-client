// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddSeries Creates a new series in the database.
type AddSeries struct {
	ApiRequest
	client ApiClient
}

// NewAddSeries creates AddSeries request.
// Creates a new series in the database.
func NewAddSeries(client ApiClient, seriesId string) *AddSeries {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &AddSeries{
		ApiRequest{
			HttpMethod:      http.MethodPut,
			Path:            fmt.Sprintf("/series/%s", seriesId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetCascadeCreate sets the cascadeCreate parameter.
// If set to `true`, the item will be created with the same ID as the series. Default is `true`.
func (r *AddSeries) SetCascadeCreate(cascadeCreate bool) *AddSeries {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddSeries) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddSeries) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
