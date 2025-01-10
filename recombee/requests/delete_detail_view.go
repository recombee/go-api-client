// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteDetailView Deletes an existing detail view uniquely specified by (`userId`, `itemId`, and `timestamp`) or all the detail views with the given `userId` and `itemId` if `timestamp` is omitted.
type DeleteDetailView struct {
	ApiRequest
	client ApiClient
}

// NewDeleteDetailView creates DeleteDetailView request.
// Deletes an existing detail view uniquely specified by (`userId`, `itemId`, and `timestamp`) or all the detail views with the given `userId` and `itemId` if `timestamp` is omitted.
func NewDeleteDetailView(client ApiClient, userId string, itemId string) *DeleteDetailView {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	return &DeleteDetailView{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/detailviews/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// Unix timestamp of the detail view. If the `timestamp` is omitted, then all the detail views with the given `userId` and `itemId` are deleted.
func (r *DeleteDetailView) SetTimestamp(timestamp timepkg.Time) *DeleteDetailView {
	r.QueryParameters["timestamp"] = timestamp.Unix()
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteDetailView) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteDetailView) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
