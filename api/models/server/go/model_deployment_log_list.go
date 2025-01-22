// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type DeploymentLogList struct {

	Object string `json:"object"`

	Data []DeploymentLog `json:"data"`
}

// AssertDeploymentLogListRequired checks if the required fields are not zero-ed
func AssertDeploymentLogListRequired(obj DeploymentLogList) error {
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
		if err := AssertDeploymentLogRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertDeploymentLogListConstraints checks if the values respects the defined constraints
func AssertDeploymentLogListConstraints(obj DeploymentLogList) error {
	for _, el := range obj.Data {
		if err := AssertDeploymentLogConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
