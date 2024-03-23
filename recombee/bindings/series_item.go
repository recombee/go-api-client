// This file is auto-generated

package bindings

// SeriesItem Binding
type SeriesItem struct {
	// `item` iff the regular item from the catalog is to be inserted, `series` iff series is inserted as the item.
	ItemType string `json:"itemType,omitempty"`
	// ID of the item iff `itemType` is `item`. ID of the series iff `itemType` is `series`.
	ItemId string `json:"itemId,omitempty"`
	// Time index used for sorting items in the series. According to time, items are sorted within series in ascending order. In the example of TV show episodes, the episode number is a natural choice to be passed as time.
	Time float64 `json:"time,omitempty"`
}
