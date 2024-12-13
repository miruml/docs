// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Miru API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1.0
 */

package openapi


import (
	"time"
)



type GitHubCommit struct {

	Object string `json:"object"`

	Sha string `json:"sha"`

	Message string `json:"message"`

	HtmlUrl string `json:"html_url"`

	PushedAt time.Time `json:"pushed_at"`

	Committer GitHubCommitter `json:"committer"`

	IsBuilt bool `json:"is_built"`

	ArtifactId *string `json:"artifact_id"`
}

// AssertGitHubCommitRequired checks if the required fields are not zero-ed
func AssertGitHubCommitRequired(obj GitHubCommit) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"sha": obj.Sha,
		"message": obj.Message,
		"html_url": obj.HtmlUrl,
		"pushed_at": obj.PushedAt,
		"committer": obj.Committer,
		"is_built": obj.IsBuilt,
		"artifact_id": obj.ArtifactId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertGitHubCommitterRequired(obj.Committer); err != nil {
		return err
	}
	return nil
}

// AssertGitHubCommitConstraints checks if the values respects the defined constraints
func AssertGitHubCommitConstraints(obj GitHubCommit) error {
	if err := AssertGitHubCommitterConstraints(obj.Committer); err != nil {
		return err
	}
	return nil
}
