// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListScenarios Get all [Scenarios](https://docs.recombee.com/scenarios) of the given database.
type ListScenarios struct {
	ApiRequest
	client ApiClient
}

// NewListScenarios creates ListScenarios request.
// Get all [Scenarios](https://docs.recombee.com/scenarios) of the given database.
func NewListScenarios(client ApiClient) *ListScenarios {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListScenarios{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/scenarios/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new([]bindings.Scenario),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *ListScenarios) SendWithContext(ctx context.Context) ([]bindings.Scenario, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Scenario)), err
}

// Sends the request to the Recombee API
func (r *ListScenarios) Send() ([]bindings.Scenario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
