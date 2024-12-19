// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type VerifyComposeFile struct {

	Object string `json:"object"`

	Content string `json:"content"`
}

// AssertVerifyComposeFileRequired checks if the required fields are not zero-ed
func AssertVerifyComposeFileRequired(obj VerifyComposeFile) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"content": obj.Content,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertVerifyComposeFileConstraints checks if the values respects the defined constraints
func AssertVerifyComposeFileConstraints(obj VerifyComposeFile) error {
	return nil
}