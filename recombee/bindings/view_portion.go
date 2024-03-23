// This file is auto-generated

package bindings

import (
	"encoding/json"
	"math"
	"time"
)

// ViewPortion Binding
type ViewPortion struct {
	// User who viewed a portion of the item
	UserId string `json:"userId,omitempty"`
	// Viewed item
	ItemId string `json:"itemId,omitempty"`
	// Viewed portion of the item (number between 0.0 (viewed nothing) and 1.0 (viewed full item) ). It should be the actual viewed part of the item, no matter the seeking. For example, if the user seeked immediately to half of the item and then viewed 10% of the item, the `portion` should still be `0.1`.
	Portion float64 `json:"portion,omitempty"`
	// ID of the session in which the user viewed the item. Default is `null` (`None`, `nil`, `NULL` etc., depending on the language).
	SessionId *string `json:"sessionId,omitempty"`
	// UTC timestamp of the rating as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// If this view portion is based on a recommendation request, `recommId` is the id of the clicked recommendation.
	RecommId *string `json:"recommId,omitempty"`
	// A dictionary of additional data for the interaction.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
}

// Support both ISO 8601 and epoch time
func (b *ViewPortion) UnmarshalJSON(data []byte) error {
	type Alias ViewPortion
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
