// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// DeleteViewPortion Deletes an existing view portion specified by (`userId`, `itemId`, `sessionId`) from the database.
type DeleteViewPortion struct {
	ApiRequest
	client ApiClient
}

// NewDeleteViewPortion creates DeleteViewPortion request.
// Deletes an existing view portion specified by (`userId`, `itemId`, `sessionId`) from the database.
func NewDeleteViewPortion(client ApiClient, userId string, itemId string) *DeleteViewPortion {

	bodyParameters := map[string]interface{}{}

	queryParams := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	return &DeleteViewPortion{
		ApiRequest{
			HttpMethod:      http.MethodDelete,
			Path:            fmt.Sprintf("/viewportions/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetSessionId sets the sessionId parameter.
// Identifier of a session.
func (r *DeleteViewPortion) SetSessionId(sessionId string) *DeleteViewPortion {
	r.QueryParameters["sessionId"] = sessionId
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *DeleteViewPortion) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *DeleteViewPortion) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
