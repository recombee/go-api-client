// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddCartAddition Adds a cart addition of the given item made by the given user.
type AddCartAddition struct {
	ApiRequest
	client ApiClient
}

// NewAddCartAddition creates AddCartAddition request.
// Adds a cart addition of the given item made by the given user.
func NewAddCartAddition(client ApiClient, userId string, itemId string) *AddCartAddition {

	bodyParameters := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	queryParams := map[string]interface{}{}

	return &AddCartAddition{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/cartadditions/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  1000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// UTC timestamp of the cart addition as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
func (r *AddCartAddition) SetTimestamp(timestamp timepkg.Time) *AddCartAddition {
	r.BodyParameters["timestamp"] = timestamp.Unix()
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the given user/item should be created if not present in the database.
func (r *AddCartAddition) SetCascadeCreate(cascadeCreate bool) *AddCartAddition {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetAmount sets the amount parameter.
// Amount (number) added to cart. The default is 1. For example, if `user-x` adds two `item-y` during a single order (session...), the `amount` should equal 2.
func (r *AddCartAddition) SetAmount(amount float64) *AddCartAddition {
	r.BodyParameters["amount"] = amount
	return r
}

// SetPrice sets the price parameter.
// Price of the added item. If `amount` is greater than 1, the sum of prices of all the items should be given.
func (r *AddCartAddition) SetPrice(price float64) *AddCartAddition {
	r.BodyParameters["price"] = price
	return r
}

// SetRecommId sets the recommId parameter.
// If this cart addition is based on a recommendation request, `recommId` is the id of the clicked recommendation.
func (r *AddCartAddition) SetRecommId(recommId string) *AddCartAddition {
	r.BodyParameters["recommId"] = recommId
	return r
}

// SetAdditionalData sets the additionalData parameter.
// A dictionary of additional data for the interaction.
func (r *AddCartAddition) SetAdditionalData(additionalData map[string]interface{}) *AddCartAddition {
	r.BodyParameters["additionalData"] = additionalData
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddCartAddition) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddCartAddition) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
