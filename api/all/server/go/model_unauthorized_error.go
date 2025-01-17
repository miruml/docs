// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type UnauthorizedError struct {

	Code string `json:"code"`

	Params map[string]interface{} `json:"params"`

	Message string `json:"message"`
}

// AssertUnauthorizedErrorRequired checks if the required fields are not zero-ed
func AssertUnauthorizedErrorRequired(obj UnauthorizedError) error {
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

	return nil
}

// AssertUnauthorizedErrorConstraints checks if the values respects the defined constraints
func AssertUnauthorizedErrorConstraints(obj UnauthorizedError) error {
	return nil
}
