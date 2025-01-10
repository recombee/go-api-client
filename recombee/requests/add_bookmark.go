// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddBookmark Adds a bookmark of the given item made by the given user.
type AddBookmark struct {
	ApiRequest
	client ApiClient
}

// NewAddBookmark creates AddBookmark request.
// Adds a bookmark of the given item made by the given user.
func NewAddBookmark(client ApiClient, userId string, itemId string) *AddBookmark {

	bodyParameters := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	queryParams := map[string]interface{}{}

	return &AddBookmark{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/bookmarks/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// UTC timestamp of the bookmark as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
func (r *AddBookmark) SetTimestamp(timestamp timepkg.Time) *AddBookmark {
	r.BodyParameters["timestamp"] = timestamp.Unix()
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the given user/item should be created if not present in the database.
func (r *AddBookmark) SetCascadeCreate(cascadeCreate bool) *AddBookmark {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetRecommId sets the recommId parameter.
// If this bookmark is based on a recommendation request, `recommId` is the id of the clicked recommendation.
func (r *AddBookmark) SetRecommId(recommId string) *AddBookmark {
	r.BodyParameters["recommId"] = recommId
	return r
}

// SetAdditionalData sets the additionalData parameter.
// A dictionary of additional data for the interaction.
func (r *AddBookmark) SetAdditionalData(additionalData map[string]interface{}) *AddBookmark {
	r.BodyParameters["additionalData"] = additionalData
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddBookmark) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddBookmark) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
