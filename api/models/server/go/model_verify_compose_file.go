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

	ComposeFile string `json:"compose_file"`

	Architecture string `json:"architecture"`
}

// AssertVerifyComposeFileRequired checks if the required fields are not zero-ed
func AssertVerifyComposeFileRequired(obj VerifyComposeFile) error {
	elements := map[string]interface{}{
		"compose_file": obj.ComposeFile,
		"architecture": obj.Architecture,
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
