// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type Error4 struct {

	Error BadRequestError `json:"error"`
}

// AssertError4Required checks if the required fields are not zero-ed
func AssertError4Required(obj Error4) error {
	elements := map[string]interface{}{
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertBadRequestErrorRequired(obj.Error); err != nil {
		return err
	}
	return nil
}

// AssertError4Constraints checks if the values respects the defined constraints
func AssertError4Constraints(obj Error4) error {
	if err := AssertBadRequestErrorConstraints(obj.Error); err != nil {
		return err
	}
	return nil
}
