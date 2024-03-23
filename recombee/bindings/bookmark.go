// This file is auto-generated

package bindings

import (
	"encoding/json"
	"math"
	"time"
)

// Bookmark Binding
type Bookmark struct {
	// User who bookmarked the item
	UserId string `json:"userId,omitempty"`
	// Bookmarked item
	ItemId string `json:"itemId,omitempty"`
	// UTC timestamp of the bookmark as ISO8601-1 pattern or UTC epoch time. The default value is the current time.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// If this bookmark is based on a recommendation request, `recommId` is the id of the clicked recommendation.
	RecommId *string `json:"recommId,omitempty"`
	// A dictionary of additional data for the interaction.
	AdditionalData *map[string]interface{} `json:"additionalData,omitempty"`
}

// Support both ISO 8601 and epoch time
func (b *Bookmark) UnmarshalJSON(data []byte) error {
	type Alias Bookmark
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
