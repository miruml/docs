// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type GitHubSource struct {

	Id string `json:"id"`

	Object string `json:"object"`

	Name string `json:"name"`

	RepositoryId int64 `json:"repository_id"`

	RepositoryName string `json:"repository_name"`

	Branch string `json:"branch"`

	BuildPath string `json:"build_path"`

	Deleted bool `json:"deleted"`
}

// AssertGitHubSourceRequired checks if the required fields are not zero-ed
func AssertGitHubSourceRequired(obj GitHubSource) error {
	elements := map[string]interface{}{
		"id": obj.Id,
		"object": obj.Object,
		"name": obj.Name,
		"repository_id": obj.RepositoryId,
		"repository_name": obj.RepositoryName,
		"branch": obj.Branch,
		"build_path": obj.BuildPath,
		"deleted": obj.Deleted,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertGitHubSourceConstraints checks if the values respects the defined constraints
func AssertGitHubSourceConstraints(obj GitHubSource) error {
	return nil
}