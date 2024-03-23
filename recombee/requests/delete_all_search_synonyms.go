// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteAllSearchSynonyms Deletes all synonyms defined in the database.
type DeleteAllSearchSynonyms struct {
	ApiRequest
	client ApiClient
}

// NewDeleteAllSearchSynonyms creates DeleteAllSearchSynonyms request.
// Deletes all synonyms defined in the database.
func NewDeleteAllSearchSynonyms(client ApiClient) *DeleteAllSearchSynonyms {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteAllSearchSynonyms{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/synonyms/items/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteAllSearchSynonyms) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteAllSearchSynonyms) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
