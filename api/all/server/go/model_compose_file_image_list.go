// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ComposeFileImageList struct {

	Object string `json:"object"`

	Data []ComposeFileImage `json:"data"`
}

// AssertComposeFileImageListRequired checks if the required fields are not zero-ed
func AssertComposeFileImageListRequired(obj ComposeFileImageList) error {
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
		if err := AssertComposeFileImageRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertComposeFileImageListConstraints checks if the values respects the defined constraints
func AssertComposeFileImageListConstraints(obj ComposeFileImageList) error {
	for _, el := range obj.Data {
		if err := AssertComposeFileImageConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
