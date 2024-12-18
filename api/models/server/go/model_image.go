// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type Image struct {

	Object string `json:"object"`

	Id string `json:"id"`

	RegistryUrl string `json:"registry_url"`

	AccountLogin string `json:"account_login"`

	Name string `json:"name"`

	Uri string `json:"uri"`

	Type string `json:"type"`

	Digest string `json:"digest"`

	Tag string `json:"tag"`
}

// AssertImageRequired checks if the required fields are not zero-ed
func AssertImageRequired(obj Image) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"registry_url": obj.RegistryUrl,
		"account_login": obj.AccountLogin,
		"name": obj.Name,
		"uri": obj.Uri,
		"type": obj.Type,
		"digest": obj.Digest,
		"tag": obj.Tag,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertImageConstraints checks if the values respects the defined constraints
func AssertImageConstraints(obj Image) error {
	return nil
}
