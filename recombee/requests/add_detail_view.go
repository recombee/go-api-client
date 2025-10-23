// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddDetailView Adds a detail view of the given item made by the given user.
type AddDetailView struct {
	ApiRequest
	client ApiClient
}

// NewAddDetailView creates AddDetailView request.
// Adds a detail view of the given item made by the given user.
func NewAddDetailView(client ApiClient, userId string, itemId string) *AddDetailView {

	bodyParameters := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	queryParams := map[string]interface{}{}

	return &AddDetailView{
		ApiRequest{
			HttpMethod:      http.MethodPost,
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
// UTC timestamp of the view as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
func (r *AddDetailView) SetTimestamp(timestamp timepkg.Time) *AddDetailView {
	r.BodyParameters["timestamp"] = timestamp.Unix()
	return r
}

// SetDuration sets the duration parameter.
// Duration of the view
func (r *AddDetailView) SetDuration(duration int) *AddDetailView {
	r.BodyParameters["duration"] = duration
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the given user/item should be created if not present in the database.
func (r *AddDetailView) SetCascadeCreate(cascadeCreate bool) *AddDetailView {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetRecommId sets the recommId parameter.
// If this detail view is based on a recommendation request, `recommId` is the id of the clicked recommendation.
func (r *AddDetailView) SetRecommId(recommId string) *AddDetailView {
	r.BodyParameters["recommId"] = recommId
	return r
}

// SetAdditionalData sets the additionalData parameter.
// A dictionary of additional data for the interaction.
func (r *AddDetailView) SetAdditionalData(additionalData map[string]interface{}) *AddDetailView {
	r.BodyParameters["additionalData"] = additionalData
	return r
}

// SetAutoPresented sets the autoPresented parameter.
// Indicates whether the item was automatically presented to the user (e.g., in a swiping feed) or explicitly requested by the user (e.g., by clicking on a link). Defaults to `false`.
func (r *AddDetailView) SetAutoPresented(autoPresented bool) *AddDetailView {
	r.BodyParameters["autoPresented"] = autoPresented
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddDetailView) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddDetailView) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
