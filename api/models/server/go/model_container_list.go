// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ContainerList struct {

	Object string `json:"object"`

	Data []Container `json:"data"`
}

// AssertContainerListRequired checks if the required fields are not zero-ed
func AssertContainerListRequired(obj ContainerList) error {
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
		if err := AssertContainerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertContainerListConstraints checks if the values respects the defined constraints
func AssertContainerListConstraints(obj ContainerList) error {
	for _, el := range obj.Data {
		if err := AssertContainerConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
