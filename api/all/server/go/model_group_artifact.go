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



type GroupArtifact struct {

	Object string `json:"object"`

	Id string `json:"id"`

	Status ArtifactStatus `json:"status"`

	Digest string `json:"digest"`

	Aarch64 bool `json:"aarch64"`

	X8664 bool `json:"x86_64"`

	CreatedAt time.Time `json:"created_at"`

	ReadyAt *time.Time `json:"ready_at"`

	FailedAt *time.Time `json:"failed_at"`

	SourceId string `json:"source_id"`

	SourceType string `json:"source_type"`

	CreatedBy *User `json:"created_by"`

	RegistrySource *RegistrySource `json:"registry_source"`

	GithubSource *GitHubSource `json:"github_source"`

	GithubSourceData *GitHubSourceData `json:"github_source_data"`

	Images ImageList `json:"images"`

	Staged bool `json:"staged"`

	Deployments GroupArtifactDeploymentList `json:"deployments"`
}

// AssertGroupArtifactRequired checks if the required fields are not zero-ed
func AssertGroupArtifactRequired(obj GroupArtifact) error {
	elements := map[string]interface{}{
		"object": obj.Object,
		"id": obj.Id,
		"status": obj.Status,
		"digest": obj.Digest,
		"aarch64": obj.Aarch64,
		"x86_64": obj.X8664,
		"created_at": obj.CreatedAt,
		"ready_at": obj.ReadyAt,
		"failed_at": obj.FailedAt,
		"source_id": obj.SourceId,
		"source_type": obj.SourceType,
		"created_by": obj.CreatedBy,
		"registry_source": obj.RegistrySource,
		"github_source": obj.GithubSource,
		"github_source_data": obj.GithubSourceData,
		"images": obj.Images,
		"staged": obj.Staged,
		"deployments": obj.Deployments,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if obj.CreatedBy != nil {
		if err := AssertUserRequired(*obj.CreatedBy); err != nil {
			return err
		}
	}
	if obj.RegistrySource != nil {
		if err := AssertRegistrySourceRequired(*obj.RegistrySource); err != nil {
			return err
		}
	}
	if obj.GithubSource != nil {
		if err := AssertGitHubSourceRequired(*obj.GithubSource); err != nil {
			return err
		}
	}
	if obj.GithubSourceData != nil {
		if err := AssertGitHubSourceDataRequired(*obj.GithubSourceData); err != nil {
			return err
		}
	}
	if err := AssertImageListRequired(obj.Images); err != nil {
		return err
	}
	if err := AssertGroupArtifactDeploymentListRequired(obj.Deployments); err != nil {
		return err
	}
	return nil
}

// AssertGroupArtifactConstraints checks if the values respects the defined constraints
func AssertGroupArtifactConstraints(obj GroupArtifact) error {
    if obj.CreatedBy != nil {
     	if err := AssertUserConstraints(*obj.CreatedBy); err != nil {
     		return err
     	}
    }
    if obj.RegistrySource != nil {
     	if err := AssertRegistrySourceConstraints(*obj.RegistrySource); err != nil {
     		return err
     	}
    }
    if obj.GithubSource != nil {
     	if err := AssertGitHubSourceConstraints(*obj.GithubSource); err != nil {
     		return err
     	}
    }
    if obj.GithubSourceData != nil {
     	if err := AssertGitHubSourceDataConstraints(*obj.GithubSourceData); err != nil {
     		return err
     	}
    }
	if err := AssertImageListConstraints(obj.Images); err != nil {
		return err
	}
	if err := AssertGroupArtifactDeploymentListConstraints(obj.Deployments); err != nil {
		return err
	}
	return nil
}