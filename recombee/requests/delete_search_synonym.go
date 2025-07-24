// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteSearchSynonym Deletes synonym of the given `id`. This synonym is no longer taken into account in the [Search items](https://docs.recombee.com/api#search-items).
type DeleteSearchSynonym struct {
	ApiRequest
	client ApiClient
}

// NewDeleteSearchSynonym creates DeleteSearchSynonym request.
// Deletes synonym of the given `id`. This synonym is no longer taken into account in the [Search items](https://docs.recombee.com/api#search-items).
func NewDeleteSearchSynonym(client ApiClient, id string) *DeleteSearchSynonym {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteSearchSynonym{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/synonyms/items/%s", id),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteSearchSynonym) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteSearchSynonym) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
