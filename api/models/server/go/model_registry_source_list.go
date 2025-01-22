// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type RegistrySourceList struct {

	Object string `json:"object"`

	Data []RegistrySource `json:"data"`
}

// AssertRegistrySourceListRequired checks if the required fields are not zero-ed
func AssertRegistrySourceListRequired(obj RegistrySourceList) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"data": obj.Data,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Data {
		if err := AssertRegistrySourceRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRegistrySourceListConstraints checks if the values respects the defined constraints
func AssertRegistrySourceListConstraints(obj RegistrySourceList) error {
	for _, el := range obj.Data {
		if err := AssertRegistrySourceConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
