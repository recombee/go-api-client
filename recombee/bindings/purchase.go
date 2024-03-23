// This file is auto-generated

package bindings

import (
	"encoding/json"
	"math"
	"time"
)

// Purchase Binding
type Purchase struct {
	// User who purchased the item
	UserId string `json:"userId,omitempty"`
	// Purchased item
	ItemId string `json:"itemId,omitempty"`
	// UTC timestamp of the purchase as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// Amount (number) of purchased items. The default is 1. For example, if `user-x` purchases two `item-y` during a single order (session...), the `amount` should equal 2.
	Amount *float64 `json:"amount,omitempty"`
	// Price paid by the user for the item. If `amount` is greater than 1, the sum of prices of all the items should be given.
	Price *float64 `json:"price,omitempty"`
	// Your profit from the purchased item. The profit is natural in the e-commerce domain (for example, if `user-x` purchases `item-y` for $100 and the gross margin is 30 %, then the profit is $30) but is also applicable in other domains (for example, at a news company it may be income from a displayed advertisement on article page). If `amount` is greater than 1, the sum of profit of all the items should be given.
	Profit *float64 `json:"profit,omitempty"`
	// If this purchase is based on a recommendation request, `recommId` is the id of the clicked recommendation.
	RecommId *string `json:"recommId,omitempty"`
	// A dictionary of additional data for the interaction.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
}

// Support both ISO 8601 and epoch time
func (b *Purchase) UnmarshalJSON(data []byte) error {
	type Alias Purchase
	temp := &struct {
		Timestamp *json.RawMessage `json:"timestamp,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(b),
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	if temp.Timestamp != nil {
		var ts time.Time
		// Attempt to unmarshal as string (ISO 8601 format)
		if err := json.Unmarshal(*temp.Timestamp, &ts); err != nil {
			// If it fails, try to unmarshal as float64 (epoch time)
			var epoch float64
			if err := json.Unmarshal(*temp.Timestamp, &epoch); err != nil {
				return err
			}
			// Convert epoch to *time.Time
			sec, dec := math.Modf(epoch)
			ts = time.Unix(int64(sec), int64(dec*(1e9)))
		}
		b.Timestamp = &ts
	}

	return nil
}
