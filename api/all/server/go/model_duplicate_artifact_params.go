// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type DuplicateArtifactParams struct {

	DuplicateArtifacts []map[string]interface{} `json:"duplicate_artifacts"`
}

// AssertDuplicateArtifactParamsRequired checks if the required fields are not zero-ed
func AssertDuplicateArtifactParamsRequired(obj DuplicateArtifactParams) error {
	elements := map[string]interface{}{
		"duplicate_artifacts": obj.DuplicateArtifacts,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertDuplicateArtifactParamsConstraints checks if the values respects the defined constraints
func AssertDuplicateArtifactParamsConstraints(obj DuplicateArtifactParams) error {
	return nil
}