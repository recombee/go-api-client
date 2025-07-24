// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteItem Deletes an item of the given `itemId` from the catalog.
// If there are any *purchases*, *ratings*, *bookmarks*, *cart additions*, or *detail views* of the item present in the database, they will be deleted in cascade as well. Also, if the item is present in some *series*, it will be removed from all the *series* where present.
// If an item becomes obsolete/no longer available, it is meaningful to keep it in the catalog (along with all the interaction data, which are very useful), and **only exclude the item from recommendations**. In such a case, use [ReQL filter](https://docs.recombee.com/reql) instead of deleting the item completely.
type DeleteItem struct {
	ApiRequest
	client ApiClient
}

// NewDeleteItem creates DeleteItem request.
// Deletes an item of the given `itemId` from the catalog.
// If there are any *purchases*, *ratings*, *bookmarks*, *cart additions*, or *detail views* of the item present in the database, they will be deleted in cascade as well. Also, if the item is present in some *series*, it will be removed from all the *series* where present.
// If an item becomes obsolete/no longer available, it is meaningful to keep it in the catalog (along with all the interaction data, which are very useful), and **only exclude the item from recommendations**. In such a case, use [ReQL filter](https://docs.recombee.com/reql) instead of deleting the item completely.
func NewDeleteItem(client ApiClient, itemId string) *DeleteItem {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{}

	return &DeleteItem{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/items/%s", itemId),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteItem) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteItem) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
