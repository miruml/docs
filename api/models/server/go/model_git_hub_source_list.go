// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type GitHubSourceList struct {

	Object string `json:"object"`

	Data []GitHubSource `json:"data"`
}

// AssertGitHubSourceListRequired checks if the required fields are not zero-ed
func AssertGitHubSourceListRequired(obj GitHubSourceList) error {
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
		if err := AssertGitHubSourceRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGitHubSourceListConstraints checks if the values respects the defined constraints
func AssertGitHubSourceListConstraints(obj GitHubSourceList) error {
	for _, el := range obj.Data {
		if err := AssertGitHubSourceConstraints(el); err != nil {
			return err
		}
	}
	return nil
}