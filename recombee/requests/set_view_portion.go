// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// SetViewPortion Sets viewed portion of an item (for example a video or article) by a user (at a session).
// If you send a new request with the same (`userId`, `itemId`, `sessionId`), the portion gets updated.
type SetViewPortion struct {
	ApiRequest
	client ApiClient
}

// NewSetViewPortion creates SetViewPortion request.
// Sets viewed portion of an item (for example a video or article) by a user (at a session).
// If you send a new request with the same (`userId`, `itemId`, `sessionId`), the portion gets updated.
func NewSetViewPortion(client ApiClient, userId string, itemId string, portion float64) *SetViewPortion {

	bodyParameters := map[string]interface{}{
		"userId":  userId,
		"itemId":  itemId,
		"portion": portion,
	}

	queryParams := map[string]interface{}{}

	return &SetViewPortion{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/viewportions/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetSessionId sets the sessionId parameter.
// ID of the session in which the user viewed the item. Default is `null` (`None`, `nil`, `NULL` etc., depending on the language).
func (r *SetViewPortion) SetSessionId(sessionId string) *SetViewPortion {
	r.BodyParameters["sessionId"] = sessionId
	return r
}

// SetTimestamp sets the timestamp parameter.
// UTC timestamp of the rating as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
func (r *SetViewPortion) SetTimestamp(timestamp timepkg.Time) *SetViewPortion {
	r.BodyParameters["timestamp"] = timestamp.Unix()
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the given user/item should be created if not present in the database.
func (r *SetViewPortion) SetCascadeCreate(cascadeCreate bool) *SetViewPortion {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetRecommId sets the recommId parameter.
// If this view portion is based on a recommendation request, `recommId` is the id of the clicked recommendation.
func (r *SetViewPortion) SetRecommId(recommId string) *SetViewPortion {
	r.BodyParameters["recommId"] = recommId
	return r
}

// SetAdditionalData sets the additionalData parameter.
// A dictionary of additional data for the interaction.
func (r *SetViewPortion) SetAdditionalData(additionalData map[string]interface{}) *SetViewPortion {
	r.BodyParameters["additionalData"] = additionalData
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *SetViewPortion) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *SetViewPortion) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
