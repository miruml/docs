/*
Miru API

Testing ArtifactsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_ArtifactsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ArtifactsAPIService ArtifactsArtifactIdDeploymentFilesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var artifactId string

		resp, httpRes, err := apiClient.ArtifactsAPI.ArtifactsArtifactIdDeploymentFilesGet(context.Background(), artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtifactsAPIService ArtifactsArtifactIdDeploymentsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var artifactId string

		resp, httpRes, err := apiClient.ArtifactsAPI.ArtifactsArtifactIdDeploymentsGet(context.Background(), artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtifactsAPIService ArtifactsArtifactIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var artifactId string

		resp, httpRes, err := apiClient.ArtifactsAPI.ArtifactsArtifactIdGet(context.Background(), artifactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtifactsAPIService GroupsGroupIdArtifactsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var groupId string

		resp, httpRes, err := apiClient.ArtifactsAPI.GroupsGroupIdArtifactsGet(context.Background(), groupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtifactsAPIService SourcesGithubGithubSourceIdArtifactsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var githubSourceId string

		httpRes, err := apiClient.ArtifactsAPI.SourcesGithubGithubSourceIdArtifactsPost(context.Background(), githubSourceId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ArtifactsAPIService SourcesRegistryRegistrySourceIdArtifactsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var registrySourceId string

		resp, httpRes, err := apiClient.ArtifactsAPI.SourcesRegistryRegistrySourceIdArtifactsPost(context.Background(), registrySourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
