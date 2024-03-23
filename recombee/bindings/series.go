// This file is auto-generated

package bindings

import (
	"encoding/json"
)

// Series Binding
type Series struct {
	// Id of the series
	SeriesId string `json:"seriesId,omitempty"`
}

func (b *Series) UnmarshalJSON(data []byte) error {
	var seriesId string
	err := json.Unmarshal(data, &seriesId)
	if err == nil {
		b.SeriesId = seriesId
		return nil
	}
	return err
}
