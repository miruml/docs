// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type Group struct {

	Object string `json:"object"`

	Id string `json:"id"`

	Name string `json:"name"`
}

// AssertGroupRequired checks if the required fields are not zero-ed
func AssertGroupRequired(obj Group) error {
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

// AssertGroupConstraints checks if the values respects the defined constraints
func AssertGroupConstraints(obj Group) error {
	return nil
}
