package bindings

import (
	"encoding/json"
)

// Item Binding
type Item struct {

	// Id of the item
	ItemId string `json:"itemId,omitempty"`

	// Property values of the item
	Values *map[string]interface{} `json:"values,omitempty"`
}

func (b *Item) UnmarshalJSON(data []byte) error {
	// Try format with values first
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err == nil {
		itemId, itemIdExists := tempMap["itemId"]
		if itemIdExists {
			// Ensure itemId is a string.
			itemIdStr, ok := itemId.(string)
			if !ok {
				return json.Unmarshal(data, &itemIdStr) // Fallback to unmarshal as string directly
			}
			b.ItemId = itemIdStr

			delete(tempMap, "itemId")

			b.Values = &tempMap
			return nil
		}
	}

	// string ID format
	var justItemId string
	err := json.Unmarshal(data, &justItemId)
	if err == nil {
		b.ItemId = justItemId
		return nil
	}
	return err
}
