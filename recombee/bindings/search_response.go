// This file is auto-generated

package bindings

// SearchResponse Binding
type SearchResponse struct {
	// Id of the personalized search request
	RecommId string `json:"recommId,omitempty"`
	// Results of the personalized search
	Recomms []Recommendation `json:"recomms,omitempty"`
	// How many times *Recommend Next Items* have been called for this `recommId`
	NumberNextRecommsCalls *int `json:"numberNextRecommsCalls,omitempty"`
	// Name of AB-testing group to which the request belongs if there is a custom AB-testing running.
	AbGroup *string `json:"abGroup,omitempty"`
}
