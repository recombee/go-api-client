// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteBookmark Deletes a bookmark uniquely specified by `userId`, `itemId`, and `timestamp` or all the bookmarks with the given `userId` and `itemId` if `timestamp` is omitted.
type DeleteBookmark struct {
	ApiRequest
	client ApiClient
}

// NewDeleteBookmark creates DeleteBookmark request.
// Deletes a bookmark uniquely specified by `userId`, `itemId`, and `timestamp` or all the bookmarks with the given `userId` and `itemId` if `timestamp` is omitted.
func NewDeleteBookmark(client ApiClient, userId string, itemId string) *DeleteBookmark {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	return &DeleteBookmark{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/bookmarks/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// Unix timestamp of the bookmark. If the `timestamp` is omitted, then all the bookmarks with the given `userId` and `itemId` are deleted.
func (r *DeleteBookmark) SetTimestamp(timestamp timepkg.Time) *DeleteBookmark {
	r.QueryParameters["timestamp"] = timestamp.Unix()
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteBookmark) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteBookmark) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
