// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type RegistrySourceContainerRepository struct {

	Object string `json:"object"`

	Id string `json:"id"`

	RegistryUrl string `json:"registry_url"`

	AccountLogin string `json:"account_login"`

	Name string `json:"name"`

	Uri string `json:"uri"`

	Type ContainerRepositoryType `json:"type"`

	IsExtra bool `json:"is_extra"`

	ComposeService *string `json:"compose_service"`
}

// AssertRegistrySourceContainerRepositoryRequired checks if the required fields are not zero-ed
func AssertRegistrySourceContainerRepositoryRequired(obj RegistrySourceContainerRepository) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"registry_url": obj.RegistryUrl,
		"account_login": obj.AccountLogin,
		"name": obj.Name,
		"uri": obj.Uri,
		"type": obj.Type,
		"is_extra": obj.IsExtra,
		"compose_service": obj.ComposeService,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRegistrySourceContainerRepositoryConstraints checks if the values respects the defined constraints
func AssertRegistrySourceContainerRepositoryConstraints(obj RegistrySourceContainerRepository) error {
	return nil
}
