// This file is auto-generated

package bindings

// GroupItem Binding
type GroupItem struct {
	// `item` iff the regular item from the catalog is to be inserted, `group` iff group is inserted as the item.
	ItemType string `json:"itemType,omitempty"`
	// ID of the item iff `itemType` is `item`. ID of the group iff `itemType` is `group`.
	ItemId string `json:"itemId,omitempty"`
}
