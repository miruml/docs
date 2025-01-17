// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type Error3 struct {

	Error ForbiddenError `json:"error"`
}

// AssertError3Required checks if the required fields are not zero-ed
func AssertError3Required(obj Error3) error {
	elements := map[string]interface{}{
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertForbiddenErrorRequired(obj.Error); err != nil {
		return err
	}
	return nil
}

// AssertError3Constraints checks if the values respects the defined constraints
func AssertError3Constraints(obj Error3) error {
	if err := AssertForbiddenErrorConstraints(obj.Error); err != nil {
		return err
	}
	return nil
}
