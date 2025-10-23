// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v6/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListSearchSynonyms Gives the list of synonyms defined in the database.
type ListSearchSynonyms struct {
	ApiRequest
	client ApiClient
}

// NewListSearchSynonyms creates ListSearchSynonyms request.
// Gives the list of synonyms defined in the database.
func NewListSearchSynonyms(client ApiClient) *ListSearchSynonyms {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListSearchSynonyms{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/synonyms/items/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(bindings.ListSearchSynonymsResponse),
		},
		client,
	}
}

// SetCount sets the count parameter.
// The number of synonyms to be listed.
func (r *ListSearchSynonyms) SetCount(count int) *ListSearchSynonyms {
	r.QueryParameters["count"] = count
	return r
}

// SetOffset sets the offset parameter.
// Specifies the number of synonyms to skip (ordered by `term`).
func (r *ListSearchSynonyms) SetOffset(offset int) *ListSearchSynonyms {
	r.QueryParameters["offset"] = offset
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *ListSearchSynonyms) SendWithContext(ctx context.Context) (bindings.ListSearchSynonymsResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.ListSearchSynonymsResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.ListSearchSynonymsResponse)), err
}

// Sends the request to the Recombee API
func (r *ListSearchSynonyms) Send() (bindings.ListSearchSynonymsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
