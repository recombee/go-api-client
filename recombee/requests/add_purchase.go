// This file is auto-generated

package requests

import (
	"context"
	"fmt"
	"net/http"
	timepkg "time" // avoid collision with param name
)

// AddPurchase Adds a purchase of the given item made by the given user.
type AddPurchase struct {
	ApiRequest
	client ApiClient
}

// NewAddPurchase creates AddPurchase request.
// Adds a purchase of the given item made by the given user.
func NewAddPurchase(client ApiClient, userId string, itemId string) *AddPurchase {

	bodyParameters := map[string]interface{}{
		"userId": userId,
		"itemId": itemId,
	}

	queryParams := map[string]interface{}{}

	return &AddPurchase{
		ApiRequest{
			HttpMethod:      http.MethodPost,
			Path:            fmt.Sprintf("/purchases/"),
			BodyParameters:  bodyParameters,
			QueryParameters: queryParams,
			DefaultTimeout:  3000 * timepkg.Millisecond,
			Target:          new(string),
		},
		client,
	}
}

// SetTimestamp sets the timestamp parameter.
// UTC timestamp of the purchase as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
func (r *AddPurchase) SetTimestamp(timestamp timepkg.Time) *AddPurchase {
	r.BodyParameters["timestamp"] = timestamp.Unix()
	return r
}

// SetCascadeCreate sets the cascadeCreate parameter.
// Sets whether the given user/item should be created if not present in the database.
func (r *AddPurchase) SetCascadeCreate(cascadeCreate bool) *AddPurchase {
	r.BodyParameters["cascadeCreate"] = cascadeCreate
	return r
}

// SetAmount sets the amount parameter.
// Amount (number) of purchased items. The default is 1. For example, if `user-x` purchases two `item-y` during a single order (session...), the `amount` should equal 2.
func (r *AddPurchase) SetAmount(amount float64) *AddPurchase {
	r.BodyParameters["amount"] = amount
	return r
}

// SetPrice sets the price parameter.
// Price paid by the user for the item. If `amount` is greater than 1, the sum of prices of all the items should be given.
func (r *AddPurchase) SetPrice(price float64) *AddPurchase {
	r.BodyParameters["price"] = price
	return r
}

// SetProfit sets the profit parameter.
// Your profit from the purchased item. The profit is natural in the e-commerce domain (for example, if `user-x` purchases `item-y` for $100 and the gross margin is 30 %, then the profit is $30) but is also applicable in other domains (for example, at a news company it may be income from a displayed advertisement on article page). If `amount` is greater than 1, the sum of profit of all the items should be given.
func (r *AddPurchase) SetProfit(profit float64) *AddPurchase {
	r.BodyParameters["profit"] = profit
	return r
}

// SetRecommId sets the recommId parameter.
// If this purchase is based on a recommendation request, `recommId` is the id of the clicked recommendation.
func (r *AddPurchase) SetRecommId(recommId string) *AddPurchase {
	r.BodyParameters["recommId"] = recommId
	return r
}

// SetAdditionalData sets the additionalData parameter.
// A dictionary of additional data for the interaction.
func (r *AddPurchase) SetAdditionalData(additionalData map[string]interface{}) *AddPurchase {
	r.BodyParameters["additionalData"] = additionalData
	return r
}

// Sends the request to the Recombee API using the specified Context
func (r *AddPurchase) SendWithContext(ctx context.Context) (string, error) {
	err := r.client.SendRequestWithContext(ctx, &r.ApiRequest)
	if err != nil {
		return "", err
	}
	return *(r.ApiRequest.Target.(*string)), err
}

// Sends the request to the Recombee API
func (r *AddPurchase) Send() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.ApiRequest.DefaultTimeout)
	defer cancel()
	return r.SendWithContext(ctx)
}
