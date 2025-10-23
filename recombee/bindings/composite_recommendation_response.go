// This file is auto-generated

package bindings

// CompositeRecommendationResponse Binding
type CompositeRecommendationResponse struct {
	// Id of the composite recommendation request
	RecommId string `json:"recommId,omitempty"`
	// Parameters of the source stage
	Source Recommendation `json:"source,omitempty"`
	// Obtained recommendations
	Recomms []Recommendation `json:"recomms,omitempty"`
	// How many times *Recommend Next Items* have been called for this `recommId`
	NumberNextRecommsCalls *int `json:"numberNextRecommsCalls,omitempty"`
}
