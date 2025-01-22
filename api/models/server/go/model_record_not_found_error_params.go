// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type RecordNotFoundErrorParams struct {

	RecordId string `json:"record_id,omitempty"`

	RecordType string `json:"record_type,omitempty"`
}

// AssertRecordNotFoundErrorParamsRequired checks if the required fields are not zero-ed
func AssertRecordNotFoundErrorParamsRequired(obj RecordNotFoundErrorParams) error {
	return nil
}

// AssertRecordNotFoundErrorParamsConstraints checks if the values respects the defined constraints
func AssertRecordNotFoundErrorParamsConstraints(obj RecordNotFoundErrorParams) error {
	return nil
}
