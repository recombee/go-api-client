package bindings

import (
	"encoding/json"
)

// User Binding
type User struct {

	// Id of the user
	UserId string `json:"userId,omitempty"`

	// Property values of the user
	Values *map[string]interface{} `json:"values,omitempty"`
}

func (b *User) UnmarshalJSON(data []byte) error {
	// Try format with values first
	var tempMap map[string]interface{}
	if err := json.Unmarshal(data, &tempMap); err == nil {
		userId, userIdExists := tempMap["userId"]
		if userIdExists {
			// Ensure itemId is a string.
			userIdStr, ok := userId.(string)
			if !ok {
				return json.Unmarshal(data, &userIdStr) // Fallback to unmarshal as string directly
			}
			b.UserId = userIdStr

			delete(tempMap, "userId")

			b.Values = &tempMap
			return nil
		}
	}

	// string ID format
	var justUserId string
	err := json.Unmarshal(data, &justUserId)
	if err == nil {
		b.UserId = justUserId
		return nil
	}
	return err
}
