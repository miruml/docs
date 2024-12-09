// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ExternalContainerRepository struct {

	RegistryUrl string `json:"registry_url"`

	Name string `json:"name"`

	Uri string `json:"uri"`

	Type ContainerRepositoryType `json:"type"`

	Object string `json:"object"`

	Description *string `json:"description"`

	// The size of the repository in bytes
	Bytes *int64 `json:"bytes"`
}

// AssertExternalContainerRepositoryRequired checks if the required fields are not zero-ed
func AssertExternalContainerRepositoryRequired(obj ExternalContainerRepository) error {
	elements := map[string]interface{}{
		"registry_url": obj.RegistryUrl,
		"name": obj.Name,
		"uri": obj.Uri,
		"type": obj.Type,
		"object": obj.Object,
		"description": obj.Description,
		"bytes": obj.Bytes,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertExternalContainerRepositoryConstraints checks if the values respects the defined constraints
func AssertExternalContainerRepositoryConstraints(obj ExternalContainerRepository) error {
	return nil
}
