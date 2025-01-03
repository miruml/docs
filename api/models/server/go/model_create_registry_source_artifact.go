// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type CreateRegistrySourceArtifact struct {

	AllowDuplicate bool `json:"allow_duplicate"`

	Images []CreateRegistrySourceArtifactImagesInner `json:"images"`
}

// AssertCreateRegistrySourceArtifactRequired checks if the required fields are not zero-ed
func AssertCreateRegistrySourceArtifactRequired(obj CreateRegistrySourceArtifact) error {
	elements := map[string]interface{}{
		"allow_duplicate": obj.AllowDuplicate,
		"images": obj.Images,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.Images {
		if err := AssertCreateRegistrySourceArtifactImagesInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertCreateRegistrySourceArtifactConstraints checks if the values respects the defined constraints
func AssertCreateRegistrySourceArtifactConstraints(obj CreateRegistrySourceArtifact) error {
	for _, el := range obj.Images {
		if err := AssertCreateRegistrySourceArtifactImagesInnerConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
