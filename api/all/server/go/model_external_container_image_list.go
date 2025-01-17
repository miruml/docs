// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ExternalContainerImageList struct {

	Object string `json:"object"`

	Data []ExternalContainerImage `json:"data"`
}

// AssertExternalContainerImageListRequired checks if the required fields are not zero-ed
func AssertExternalContainerImageListRequired(obj ExternalContainerImageList) error {
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
		if err := AssertExternalContainerImageRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertExternalContainerImageListConstraints checks if the values respects the defined constraints
func AssertExternalContainerImageListConstraints(obj ExternalContainerImageList) error {
	for _, el := range obj.Data {
		if err := AssertExternalContainerImageConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
