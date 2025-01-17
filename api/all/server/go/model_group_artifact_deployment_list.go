// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type GroupArtifactDeploymentList struct {

	Object string `json:"object"`

	Data []GroupArtifactDeployment `json:"data"`
}

// AssertGroupArtifactDeploymentListRequired checks if the required fields are not zero-ed
func AssertGroupArtifactDeploymentListRequired(obj GroupArtifactDeploymentList) error {
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
		if err := AssertGroupArtifactDeploymentRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGroupArtifactDeploymentListConstraints checks if the values respects the defined constraints
func AssertGroupArtifactDeploymentListConstraints(obj GroupArtifactDeploymentList) error {
	for _, el := range obj.Data {
		if err := AssertGroupArtifactDeploymentConstraints(el); err != nil {
			return err
		}
	}
	return nil
}