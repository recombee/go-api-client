// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteCartAddition Deletes an existing cart addition uniquely specified by `userId`, `itemId`, and `timestamp` or all the cart additions with the given `userId` and `itemId` if `timestamp` is omitted.
type DeleteCartAddition struct {
	ApiRequest
	client ApiClient
}

// NewDeleteCartAddition creates DeleteCartAddition request.
// Deletes an existing cart addition uniquely specified by `userId`, `itemId`, and `timestamp` or all the cart additions with the given `userId` and `itemId` if `timestamp` is omitted.
func NewDeleteCartAddition(client ApiClient, userId string, itemId string) *DeleteCartAddition {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	return &DeleteCartAddition{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/cartadditions/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// Unix timestamp of the cart addition. If the `timestamp` is omitted, then all the cart additions with the given `userId` and `itemId` are deleted.
func (r *DeleteCartAddition) SetTimestamp(timestamp timepkg.Time) *DeleteCartAddition {
	r.QueryParameters["timestamp"] = timestamp.Unix()
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteCartAddition) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteCartAddition) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
