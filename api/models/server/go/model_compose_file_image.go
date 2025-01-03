// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type ComposeFileImage struct {

	Object string `json:"object"`

	ComposeReference string `json:"compose_reference"`

	ImageUri string `json:"image_uri"`

	RepositoryUri string `json:"repository_uri"`

	RegistryUrl string `json:"registry_url"`

	RegistryType string `json:"registry_type"`

	Name string `json:"name"`

	Digest string `json:"digest"`

	Tag string `json:"tag"`

	IsImageValid bool `json:"is_image_valid"`

	IsRepositoryValid bool `json:"is_repository_valid"`

	Error string `json:"error"`
}

// AssertComposeFileImageRequired checks if the required fields are not zero-ed
func AssertComposeFileImageRequired(obj ComposeFileImage) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"compose_reference": obj.ComposeReference,
		"image_uri": obj.ImageUri,
		"repository_uri": obj.RepositoryUri,
		"registry_url": obj.RegistryUrl,
		"registry_type": obj.RegistryType,
		"name": obj.Name,
		"digest": obj.Digest,
		"tag": obj.Tag,
		"is_image_valid": obj.IsImageValid,
		"is_repository_valid": obj.IsRepositoryValid,
		"error": obj.Error,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertComposeFileImageConstraints checks if the values respects the defined constraints
func AssertComposeFileImageConstraints(obj ComposeFileImage) error {
	return nil
}
