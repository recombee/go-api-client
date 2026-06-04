// This file is auto-generated

package bindings

// PropertyInfo Binding
type PropertyInfo struct {
	// Name of the property
	Name string `json:"name,omitempty"`
	// Type of the property
	Type string `json:"type,omitempty"`
	// [Role](https://docs.recombee.com/api/property_roles_metadata#roles) assigned to the property.
	// Role defines how the property is interpreted by the system.
	Role PropertyRole `json:"role,omitempty"`
	// [Metadata](https://docs.recombee.com/api/property_roles_metadata#metadata) entries associated with the property.
	// Metadata provides additional configuration or annotations that influence how the property is processed.
	Metadata []PropertyMetadata `json:"metadata,omitempty"`
}
