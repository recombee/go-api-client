// This file is auto-generated

package bindings

// Segmentation Binding
type Segmentation struct {
	// Id of the Segmentation
	SegmentationId string `json:"segmentationId,omitempty"`
	// Determines on which type of data (e.g. `items`) the Segmentation is based
	SourceType string `json:"sourceType,omitempty"`
	// Determines the type of the segmentation (Property-based, Manual ReQL, Auto ReQL)
	SegmentationType string `json:"segmentationType,omitempty"`
	// Title of the Segmentation
	Title *string `json:"title,omitempty"`
	// Description of the Segmentation
	Description *string `json:"description,omitempty"`
}
