// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeletePurchase Deletes an existing purchase uniquely specified by `userId`, `itemId`, and `timestamp` or all the purchases with the given `userId` and `itemId` if `timestamp` is omitted.
type DeletePurchase struct {
	ApiRequest
	client ApiClient
}

// NewDeletePurchase creates DeletePurchase request.
// Deletes an existing purchase uniquely specified by `userId`, `itemId`, and `timestamp` or all the purchases with the given `userId` and `itemId` if `timestamp` is omitted.
func NewDeletePurchase(client ApiClient, userId string, itemId string) *DeletePurchase {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	return &DeletePurchase{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/purchases/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// Unix timestamp of the purchase. If the `timestamp` is omitted, then all the purchases with the given `userId` and `itemId` are deleted.
func (r *DeletePurchase) SetTimestamp(timestamp timepkg.Time) *DeletePurchase {
	r.QueryParameters["timestamp"] = timestamp.Unix()
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *DeletePurchase) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeletePurchase) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
