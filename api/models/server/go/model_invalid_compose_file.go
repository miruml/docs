// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type InvalidComposeFile struct {

	Code string `json:"code"`

	Params InvalidComposeFileParams `json:"params"`

	Message string `json:"message"`
}

// AssertInvalidComposeFileRequired checks if the required fields are not zero-ed
func AssertInvalidComposeFileRequired(obj InvalidComposeFile) error {
	elements := map[string]interface{}{
		"code": obj.Code,
		"params": obj.Params,
		"message": obj.Message,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertInvalidComposeFileParamsRequired(obj.Params); err != nil {
		return err
	}
	return nil
}

// AssertInvalidComposeFileConstraints checks if the values respects the defined constraints
func AssertInvalidComposeFileConstraints(obj InvalidComposeFile) error {
	if err := AssertInvalidComposeFileParamsConstraints(obj.Params); err != nil {
		return err
	}
	return nil
}
