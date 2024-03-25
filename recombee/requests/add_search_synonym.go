// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v4/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddSearchSynonym Adds a new synonym for the [Search items](https://docs.recombee.com/api.html#search-items).
// When the `term` is used in the search query, the `synonym` is also used for the full-text search.
// Unless `oneWay=true`, it works also in the opposite way (`synonym` -> `term`).
// An example of a synonym can be `science fiction` for the term `sci-fi`.
type AddSearchSynonym struct {
	ApiRequest
	client ApiClient
}

// NewAddSearchSynonym creates AddSearchSynonym request.
// Adds a new synonym for the [Search items](https://docs.recombee.com/api.html#search-items).
// When the `term` is used in the search query, the `synonym` is also used for the full-text search.
// Unless `oneWay=true`, it works also in the opposite way (`synonym` -> `term`).
// An example of a synonym can be `science fiction` for the term `sci-fi`.
func NewAddSearchSynonym(client ApiClient, term string, synonym string) *AddSearchSynonym {

	bodyParameters := map[string]interface{}{
		"term":    term,
		"synonym": synonym,
	}

	queryParams := map[string]interface{}{}

	return &AddSearchSynonym{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/synonyms/items/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  10000 * timepkg.Millisecond,
			Target:          new(bindings.SearchSynonym),
		},
		client,
	}
}

// SetOneWay sets the oneWay parameter.
// If set to `true`, only `term` -> `synonym` is considered. If set to `false`, also `synonym` -> `term` works.
// Default: `false`.
func (r *AddSearchSynonym) SetOneWay(oneWay bool) *AddSearchSynonym {
	r.BodyParameters["oneWay"] = oneWay
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddSearchSynonym) SendWithContext(ctx context.Context) (bindings.SearchSynonym, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.SearchSynonym{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.SearchSynonym)), err
}

// Sends the request to the Recombee API
func (r *AddSearchSynonym) Send() (bindings.SearchSynonym, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
