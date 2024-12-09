// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ContainerRepositoryWithIsExtra struct {

	RegistryUrl string `json:"registry_url"`

	Name string `json:"name"`

	Uri string `json:"uri"`

	Type ContainerRepositoryType `json:"type"`

	Object string `json:"object"`

	Id string `json:"id"`

	IsExtra bool `json:"is_extra"`
}

// AssertContainerRepositoryWithIsExtraRequired checks if the required fields are not zero-ed
func AssertContainerRepositoryWithIsExtraRequired(obj ContainerRepositoryWithIsExtra) error {
	elements := map[string]interface{}{
		"registry_url": obj.RegistryUrl,
		"name": obj.Name,
		"uri": obj.Uri,
		"type": obj.Type,
		"object": obj.Object,
		"id": obj.Id,
		"is_extra": obj.IsExtra,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertContainerRepositoryWithIsExtraConstraints checks if the values respects the defined constraints
func AssertContainerRepositoryWithIsExtraConstraints(obj ContainerRepositoryWithIsExtra) error {
	return nil
}
