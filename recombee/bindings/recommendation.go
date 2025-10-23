// This file is auto-generated

package bindings

// Recommendation Binding
type Recommendation struct {
	// Id of the recommended item
	Id string `json:"id,omitempty"`
	// Property values of the recommended item
	Values *map[string]interface{} `json:"values,omitempty"`
	// Dictionary of evaluated ReQL expressions specified in the request and calculated for the recommended item.
	// The keys are the names of the ReQL expressions, and the values are the results of the evaluations.
	ReqlEvaluations *map[string]interface{} `json:"reqlEvaluations,omitempty"`
}
