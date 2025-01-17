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
	"context"
	"net/http"
	"errors"
)

// GitHubSourcesAPIService is a service that implements the logic for the GitHubSourcesAPIServicer
// This service should implement the business logic for every endpoint for the GitHubSourcesAPI API.
// Include any external packages or services that will be required by this service.
type GitHubSourcesAPIService struct {
}

// NewGitHubSourcesAPIService creates a default api service
func NewGitHubSourcesAPIService() *GitHubSourcesAPIService {
	return &GitHubSourcesAPIService{}
}

// SourcesGithubGithubSourceIdCommitsGet - 
func (s *GitHubSourcesAPIService) SourcesGithubGithubSourceIdCommitsGet(ctx context.Context, workspaceId string, githubSourceId string) (ImplResponse, error) {
	// TODO - update SourcesGithubGithubSourceIdCommitsGet with the required logic for this service method.
	// Add api_git_hub_sources_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GitHubSourceCommits{}) or use other options such as http.Ok ...
	// return Response(200, GitHubSourceCommits{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("SourcesGithubGithubSourceIdCommitsGet method not implemented")
}