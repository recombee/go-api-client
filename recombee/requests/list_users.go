// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// ListUsers Gets a list of IDs of users currently present in the catalog.
type ListUsers struct {
	ApiRequest
	client ApiClient
}

// NewListUsers creates ListUsers request.
// Gets a list of IDs of users currently present in the catalog.
func NewListUsers(client ApiClient) *ListUsers {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &ListUsers{
		ApiRequest{
			HttpMethod:      http.MethodGet,
			Path:            fmt.Sprintf("/users/list/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new([]bindings.User),
		},
		client,
	}
}

// SetFilter sets the filter parameter.
// Boolean-returning [ReQL](https://docs.recombee.com/reql.html) expression, which allows you to filter users to be listed. Only the users for which the expression is *true* will be returned.
func (r *ListUsers) SetFilter(filter string) *ListUsers {
	r.QueryParameters["filter"] = filter
	return r
}

// SetCount sets the count parameter.
// The number of users to be listed.
func (r *ListUsers) SetCount(count int) *ListUsers {
	r.QueryParameters["count"] = count
	return r
}

// SetOffset sets the offset parameter.
// Specifies the number of users to skip (ordered by `userId`).
func (r *ListUsers) SetOffset(offset int) *ListUsers {
	r.QueryParameters["offset"] = offset
	return r
}

// SetReturnProperties sets the returnProperties parameter.
// With `returnProperties=true`, property values of the listed users are returned along with their IDs in a JSON dictionary.
// Example response:
// ```
//
//	[
//	  {
//	    "userId": "user-81",
//	    "country": "US",
//	    "sex": "M"
//	  },
//	  {
//	    "userId": "user-314",
//	    "country": "CAN",
//	    "sex": "F"
//	  }
//	]
//
// ```
func (r *ListUsers) SetReturnProperties(returnProperties bool) *ListUsers {
	r.QueryParameters["returnProperties"] = returnProperties
	return r
}

// SetIncludedProperties sets the includedProperties parameter.
// Allows specifying which properties should be returned when `returnProperties=true` is set. The properties are given as a comma-separated list.
// Example response for `includedProperties=country`:
// ```
//
//	[
//	  {
//	    "userId": "user-81",
//	    "country": "US"
//	  },
//	  {
//	    "userId": "user-314",
//	    "country": "CAN"
//	  }
//	]
//
// ```
func (r *ListUsers) SetIncludedProperties(includedProperties []string) *ListUsers {
	r.QueryParameters["includedProperties"] = includedProperties
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *ListUsers) SendWithContext(ctx context.Context) ([]bindings.User, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return nil, err
	}
	return *(r.ApiRequest.Target.(*[]bindings.User)), err
}

// Sends the request to the Recombee API
func (r *ListUsers) Send() ([]bindings.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
