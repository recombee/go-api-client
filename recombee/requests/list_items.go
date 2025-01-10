// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListItems Gets a list of IDs of items currently present in the catalog.
type ListItems struct {
	ApiRequest
	client ApiClient
}

// NewListItems creates ListItems request.
// Gets a list of IDs of items currently present in the catalog.
func NewListItems(client ApiClient) *ListItems {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListItems{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/items/list/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.Item),
		},
		client,
	}
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to filter items to be listed. Only the items for which the expression is *true* will be returned.
func (r *ListItems) SetFilter(filter string) *ListItems {
	r.QueryParameters["filter"] = filter
	return r
}

// SetCount sets the count parameter.
// The number of items to be listed.
func (r *ListItems) SetCount(count int) *ListItems {
	r.QueryParameters["count"] = count
	return r
}

// SetOffset sets the offset parameter.
// Specifies the number of items to skip (ordered by `itemId`).
func (r *ListItems) SetOffset(offset int) *ListItems {
	r.QueryParameters["offset"] = offset
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the listed items are returned along with their IDs in a JSON dictionary.
// Example response:
// ```
//
//	[
//	  {
//	    "itemId": "tv-178",
//	    "description": "4K TV with 3D feature",
//	    "categories":   ["Electronics", "Televisions"],
//	    "price": 342,
//	    "url": "myshop.com/tv-178"
//	  },
//	  {
//	    "itemId": "mixer-42",
//	    "description": "Stainless Steel Mixer",
//	    "categories":   ["Home & Kitchen"],
//	    "price": 39,
//	    "url": "myshop.com/mixer-42"
//	  }
//	]
//
// ```
func (r *ListItems) SetReturnProperties(returnProperties bool) *ListItems {
	r.QueryParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=description,price`:
// ```
//
//	[
//	  {
//	    "itemId": "tv-178",
//	    "description": "4K TV with 3D feature",
//	    "price": 342
//	  },
//	  {
//	    "itemId": "mixer-42",
//	    "description": "Stainless Steel Mixer",
//	    "price": 39
//	  }
//	]
//
// ```
func (r *ListItems) SetIncludedProperties(includedProperties []string) *ListItems {
	r.QueryParameters["includedProperties"] = includedProperties
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *ListItems) SendWithContext(ctx context.Context) ([]bindings.Item, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.Item)), err
}

// Sends the request to the Recombee API
func (r *ListItems) Send() ([]bindings.Item, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
