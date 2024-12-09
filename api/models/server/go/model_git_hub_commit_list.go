// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi




type GitHubCommitList struct {

	Object string `json:"object,omitempty"`

	Data []GitHubCommit `json:"data,omitempty"`
}

// AssertGitHubCommitListRequired checks if the required fields are not zero-ed
func AssertGitHubCommitListRequired(obj GitHubCommitList) error {
	for _, el := range obj.Data {
		if err := AssertGitHubCommitRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertGitHubCommitListConstraints checks if the values respects the defined constraints
func AssertGitHubCommitListConstraints(obj GitHubCommitList) error {
	for _, el := range obj.Data {
		if err := AssertGitHubCommitConstraints(el); err != nil {
			return err
		}
	}
	return nil
}
