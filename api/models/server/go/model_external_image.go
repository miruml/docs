// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ExternalImage struct {

	Object string `json:"object,omitempty"`

	Digest string `json:"digest"`

	Tags []string `json:"tags"`

	Uri string `json:"uri"`

	Bytes int32 `json:"bytes"`

	UploadedAt string `json:"uploaded_at"`
}

// AssertExternalImageRequired checks if the required fields are not zero-ed
func AssertExternalImageRequired(obj ExternalImage) error {
	elements := map[string]interface{}{
		"digest": obj.Digest,
		"tags": obj.Tags,
		"uri": obj.Uri,
		"bytes": obj.Bytes,
		"uploaded_at": obj.UploadedAt,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertExternalImageConstraints checks if the values respects the defined constraints
func AssertExternalImageConstraints(obj ExternalImage) error {
	return nil
}
