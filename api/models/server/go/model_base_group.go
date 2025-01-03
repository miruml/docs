// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type BaseGroup struct {

	Object string `json:"object"`

	Id string `json:"id"`

	WorkspaceId string `json:"workspace_id,omitempty"`

	Name string `json:"name"`
}

// AssertBaseGroupRequired checks if the required fields are not zero-ed
func AssertBaseGroupRequired(obj BaseGroup) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertBaseGroupConstraints checks if the values respects the defined constraints
func AssertBaseGroupConstraints(obj BaseGroup) error {
	return nil
}
