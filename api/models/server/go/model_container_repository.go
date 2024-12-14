// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ContainerRepository struct {

	Object string `json:"object"`

	Id string `json:"id"`

	RegistryUrl string `json:"registry_url"`

	AccountLogin string `json:"account_login"`

	Name string `json:"name"`

	Uri string `json:"uri"`

	Type ContainerRepositoryType `json:"type"`
}

// AssertContainerRepositoryRequired checks if the required fields are not zero-ed
func AssertContainerRepositoryRequired(obj ContainerRepository) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"registry_url": obj.RegistryUrl,
		"account_login": obj.AccountLogin,
		"name": obj.Name,
		"uri": obj.Uri,
		"type": obj.Type,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertContainerRepositoryConstraints checks if the values respects the defined constraints
func AssertContainerRepositoryConstraints(obj ContainerRepository) error {
	return nil
}
