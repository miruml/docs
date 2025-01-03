// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type Group struct {

	Object string `json:"object"`

	Id string `json:"id"`

	WorkspaceId string `json:"workspace_id,omitempty"`

	Name string `json:"name"`

	Devices GroupDeviceList `json:"devices,omitempty"`

	GithubSources GitHubSourceList `json:"github_sources,omitempty"`
}

// AssertGroupRequired checks if the required fields are not zero-ed
func AssertGroupRequired(obj Group) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGroupDeviceListRequired(obj.Devices); err != nil {
		return err
	}
	if err := AssertGitHubSourceListRequired(obj.GithubSources); err != nil {
		return err
	}
	return nil
}

// AssertGroupConstraints checks if the values respects the defined constraints
func AssertGroupConstraints(obj Group) error {
	if err := AssertGroupDeviceListConstraints(obj.Devices); err != nil {
		return err
	}
	if err := AssertGitHubSourceListConstraints(obj.GithubSources); err != nil {
		return err
	}
	return nil
}
