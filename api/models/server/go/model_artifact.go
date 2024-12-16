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



type Artifact struct {

	Object string `json:"object"`

	Id string `json:"id"`

	Status ArtifactStatus `json:"status"`

	Aarch64 bool `json:"aarch64"`

	X8664 bool `json:"x86_64"`

	CreatedAt time.Time `json:"created_at"`

	ReadyAt *time.Time `json:"ready_at"`

	FailedAt *time.Time `json:"failed_at"`

	CreatedBy User `json:"created_by,omitempty"`

	Deployments ArtifactDeploymentList `json:"deployments,omitempty"`

	Images ImageList `json:"images"`

	SourceId string `json:"source_id"`

	SourceType string `json:"source_type"`

	RegistrySource RegistrySource `json:"registry_source,omitempty"`

	GithubSource GitHubSource `json:"github_source,omitempty"`

	GithubSourceData GitHubSourceData `json:"github_source_data,omitempty"`
}

// AssertArtifactRequired checks if the required fields are not zero-ed
func AssertArtifactRequired(obj Artifact) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"status": obj.Status,
		"aarch64": obj.Aarch64,
		"x86_64": obj.X8664,
		"created_at": obj.CreatedAt,
		"ready_at": obj.ReadyAt,
		"failed_at": obj.FailedAt,
		"images": obj.Images,
		"source_id": obj.SourceId,
		"source_type": obj.SourceType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertUserRequired(obj.CreatedBy); err != nil {
		return err
	}
	if err := AssertArtifactDeploymentListRequired(obj.Deployments); err != nil {
		return err
	}
	if err := AssertImageListRequired(obj.Images); err != nil {
		return err
	}
	if err := AssertRegistrySourceRequired(obj.RegistrySource); err != nil {
		return err
	}
	if err := AssertGitHubSourceRequired(obj.GithubSource); err != nil {
		return err
	}
	if err := AssertGitHubSourceDataRequired(obj.GithubSourceData); err != nil {
		return err
	}
	return nil
}

// AssertArtifactConstraints checks if the values respects the defined constraints
func AssertArtifactConstraints(obj Artifact) error {
	if err := AssertUserConstraints(obj.CreatedBy); err != nil {
		return err
	}
	if err := AssertArtifactDeploymentListConstraints(obj.Deployments); err != nil {
		return err
	}
	if err := AssertImageListConstraints(obj.Images); err != nil {
		return err
	}
	if err := AssertRegistrySourceConstraints(obj.RegistrySource); err != nil {
		return err
	}
	if err := AssertGitHubSourceConstraints(obj.GithubSource); err != nil {
		return err
	}
	if err := AssertGitHubSourceDataConstraints(obj.GithubSourceData); err != nil {
		return err
	}
	return nil
}
