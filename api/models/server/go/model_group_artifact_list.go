// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type GroupArtifactList struct {

	Object string `json:"object"`

	Data []GroupArtifact `json:"data"`
}

// AssertGroupArtifactListRequired checks if the required fields are not zero-ed
func AssertGroupArtifactListRequired(obj GroupArtifactList) error {
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
		if err := AssertGroupArtifactRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGroupArtifactListConstraints checks if the values respects the defined constraints
func AssertGroupArtifactListConstraints(obj GroupArtifactList) error {
	for _, el := range obj.Data {
		if err := AssertGroupArtifactConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
