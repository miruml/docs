// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ExternalContainerRepositoryList struct {

	Object string `json:"object"`

	Data []ExternalContainerRepository `json:"data"`
}

// AssertExternalContainerRepositoryListRequired checks if the required fields are not zero-ed
func AssertExternalContainerRepositoryListRequired(obj ExternalContainerRepositoryList) error {
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
		if err := AssertExternalContainerRepositoryRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertExternalContainerRepositoryListConstraints checks if the values respects the defined constraints
func AssertExternalContainerRepositoryListConstraints(obj ExternalContainerRepositoryList) error {
	for _, el := range obj.Data {
		if err := AssertExternalContainerRepositoryConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
