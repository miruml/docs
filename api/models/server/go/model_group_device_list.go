// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type GroupDeviceList struct {

	Object string `json:"object"`

	Data []GroupDevice `json:"data"`
}

// AssertGroupDeviceListRequired checks if the required fields are not zero-ed
func AssertGroupDeviceListRequired(obj GroupDeviceList) error {
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
		if err := AssertGroupDeviceRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGroupDeviceListConstraints checks if the values respects the defined constraints
func AssertGroupDeviceListConstraints(obj GroupDeviceList) error {
	for _, el := range obj.Data {
		if err := AssertGroupDeviceConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
