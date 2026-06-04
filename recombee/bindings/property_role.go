// This file is auto-generated

package bindings

// PropertyRole Binding
type PropertyRole struct {
	// Name of the [role](https://docs.recombee.com/api/property_roles_metadata#roles) assigned to the property.
	Name string `json:"name,omitempty"`
	// Optional configuration specific to the selected role.
	Settings *map[string]interface{} `json:"settings,omitempty"`
}
