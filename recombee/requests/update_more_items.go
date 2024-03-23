// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// UpdateMoreItems Updates (some) property values of all the items that pass the filter.
// Example: *Setting all the items that are older than a week as unavailable*
//
//	```
//	  {
//	    "filter": "'releaseDate' < now() - 7*24*3600",
//	    "changes": {"available": false}
//	  }
//	```
type UpdateMoreItems struct {
	ApiRequest
	client ApiClient
}

// NewUpdateMoreItems creates UpdateMoreItems request.
// Updates (some) property values of all the items that pass the filter.
// Example: *Setting all the items that are older than a week as unavailable*
//
//	```
//	  {
//	    "filter": "'releaseDate' < now() - 7*24*3600",
//	    "changes": {"available": false}
//	  }
//	```
func NewUpdateMoreItems(client ApiClient, filter string, changes map[string]interface{}) *UpdateMoreItems {

	bodyParameters := map[string]interface{}{
		"filter":  filter,
		"changes": changes,
	}

	queryParams := map[string]interface{}{}

	return &UpdateMoreItems{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/more-items/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(bindings.UpdateMoreItemsResponse),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *UpdateMoreItems) SendWithContext(ctx context.Context) (bindings.UpdateMoreItemsResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.UpdateMoreItemsResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.UpdateMoreItemsResponse)), err
}

// Sends the request to the Recombee API
func (r *UpdateMoreItems) Send() (bindings.UpdateMoreItemsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
