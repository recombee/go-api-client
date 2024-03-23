// This file is auto-generated

package bindings

// SearchSynonym Binding
type SearchSynonym struct {
	// Id of the synonym record
	Id string `json:"id,omitempty"`
	// A word to which the `synonym` is specified.
	Term string `json:"term,omitempty"`
	// A word that should be considered equal to `term` by the full-text search engine.
	Synonym string `json:"synonym,omitempty"`
	// If set to `true`, only `term` -> `synonym` is considered. I set to `false`, also `synonym` -> `term` works.
	OneWay bool `json:"oneWay,omitempty"`
}
