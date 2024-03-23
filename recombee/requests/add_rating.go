// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddRating Adds a rating of the given item made by the given user.
type AddRating struct {
	ApiRequest
	client ApiClient
}

// NewAddRating creates AddRating request.
// Adds a rating of the given item made by the given user.
func NewAddRating(client ApiClient, userId string, itemId string, rating float64) *AddRating {

	bodyParameters := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
		"rating": rating,
	}

	queryParams := map[string]interface{}{}

	return &AddRating{
		ApiRequest{
			HttpMethod:      http.MethodPost,
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
// UTC timestamp of the rating as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
func (r *AddRating) SetTimestamp(timestamp timepkg.Time) *AddRating {
	r.BodyParameters["timestamp"] = timestamp.Unix()
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the given user/item should be created if not present in the database.
func (r *AddRating) SetCascadeCreate(cascadeCreate bool) *AddRating {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetRecommId sets the recommId parameter.
// If this rating is based on a recommendation request, `recommId` is the id of the clicked recommendation.
func (r *AddRating) SetRecommId(recommId string) *AddRating {
	r.BodyParameters["recommId"] = recommId
	return r
}

// SetAdditionalData sets the additionalData parameter.
// A dictionary of additional data for the interaction.
func (r *AddRating) SetAdditionalData(additionalData map[string]interface{}) *AddRating {
	r.BodyParameters["additionalData"] = additionalData
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddRating) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddRating) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
