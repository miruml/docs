// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi


import (
	"time"
)



type DeploymentLog struct {

	Timestamp time.Time `json:"timestamp"`

	Level string `json:"level"`

	Message string `json:"message"`

	Code string `json:"code"`
}

// AssertDeploymentLogRequired checks if the required fields are not zero-ed
func AssertDeploymentLogRequired(obj DeploymentLog) error {
	elements := map[string]interface{}{
		"timestamp": obj.Timestamp,
		"level": obj.Level,
		"message": obj.Message,
		"code": obj.Code,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertDeploymentLogConstraints checks if the values respects the defined constraints
func AssertDeploymentLogConstraints(obj DeploymentLog) error {
	return nil
}
