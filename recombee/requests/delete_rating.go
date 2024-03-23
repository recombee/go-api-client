// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteRating Deletes an existing rating specified by (`userId`, `itemId`, `timestamp`) from the database or all the ratings with the given `userId` and `itemId` if `timestamp` is omitted.
type DeleteRating struct {
	ApiRequest
	client ApiClient
}

// NewDeleteRating creates DeleteRating request.
// Deletes an existing rating specified by (`userId`, `itemId`, `timestamp`) from the database or all the ratings with the given `userId` and `itemId` if `timestamp` is omitted.
func NewDeleteRating(client ApiClient, userId string, itemId string) *DeleteRating {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	return &DeleteRating{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/ratings/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// Unix timestamp of the rating. If the `timestamp` is omitted, then all the ratings with the given `userId` and `itemId` are deleted.
func (r *DeleteRating) SetTimestamp(timestamp timepkg.Time) *DeleteRating {
	r.QueryParameters["timestamp"] = timestamp.Unix()
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteRating) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteRating) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
