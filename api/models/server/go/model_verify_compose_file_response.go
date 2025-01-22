// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type VerifyComposeFileResponse struct {

	Object string `json:"object"`

	Content string `json:"content"`

	IsValid bool `json:"is_valid"`

	IsSchemaValid bool `json:"is_schema_valid"`

	SchemaError string `json:"schema_error,omitempty"`

	Images ComposeFileImageList `json:"images"`
}

// AssertVerifyComposeFileResponseRequired checks if the required fields are not zero-ed
func AssertVerifyComposeFileResponseRequired(obj VerifyComposeFileResponse) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"content": obj.Content,
		"is_valid": obj.IsValid,
		"is_schema_valid": obj.IsSchemaValid,
		"images": obj.Images,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertComposeFileImageListRequired(obj.Images); err != nil {
		return err
	}
	return nil
}

// AssertVerifyComposeFileResponseConstraints checks if the values respects the defined constraints
func AssertVerifyComposeFileResponseConstraints(obj VerifyComposeFileResponse) error {
	if err := AssertComposeFileImageListConstraints(obj.Images); err != nil {
		return err
	}
	return nil
}
