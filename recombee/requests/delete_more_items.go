// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"github.com/recombee/go-api-client/v5/recombee/bindings"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteMoreItems Deletes all the items that pass the filter.
// If an item becomes obsolete/no longer available, it is meaningful to **keep it in the catalog** (along with all the interaction data, which are very useful) and **only exclude the item from recommendations**. In such a case, use [ReQL filter](https://docs.recombee.com/reql.html) instead of deleting the item completely.
type DeleteMoreItems struct {
	ApiRequest
	client ApiClient
}

// NewDeleteMoreItems creates DeleteMoreItems request.
// Deletes all the items that pass the filter.
// If an item becomes obsolete/no longer available, it is meaningful to **keep it in the catalog** (along with all the interaction data, which are very useful) and **only exclude the item from recommendations**. In such a case, use [ReQL filter](https://docs.recombee.com/reql.html) instead of deleting the item completely.
func NewDeleteMoreItems(client ApiClient, filter string) *DeleteMoreItems {

	bodyParameters := map[string]interface{}{
		"filter": filter,
	}

	queryParams := map[string]interface{}{}

	return &DeleteMoreItems{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/more-items/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  100000 * timepkg.Millisecond,
			Target:          new(bindings.DeleteMoreItemsResponse),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteMoreItems) SendWithContext(ctx context.Context) (bindings.DeleteMoreItemsResponse, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return bindings.DeleteMoreItemsResponse{}, err
	}
	return *(r.ApiRequest.Target.(*bindings.DeleteMoreItemsResponse)), err
}

// Sends the request to the Recombee API
func (r *DeleteMoreItems) Send() (bindings.DeleteMoreItemsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
