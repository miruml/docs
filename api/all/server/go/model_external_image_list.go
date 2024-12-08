// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ExternalImageList struct {

	Object string `json:"object,omitempty"`

	Data []ExternalImage `json:"data,omitempty"`
}

// AssertExternalImageListRequired checks if the required fields are not zero-ed
func AssertExternalImageListRequired(obj ExternalImageList) error {
	for _, el := range obj.Data {
		if err := AssertExternalImageRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertExternalImageListConstraints checks if the values respects the defined constraints
func AssertExternalImageListConstraints(obj ExternalImageList) error {
	for _, el := range obj.Data {
		if err := AssertExternalImageConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
