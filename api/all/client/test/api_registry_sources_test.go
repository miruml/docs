/*
Miru API

Testing RegistrySourcesAPIService

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

func Test_openapi_RegistrySourcesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RegistrySourcesAPIService SourcesRegistryRegistrySourceIdComposeFileGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var registrySourceId string

		resp, httpRes, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdComposeFileGet(context.Background(), registrySourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrySourcesAPIService SourcesRegistryRegistrySourceIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var registrySourceId string

		resp, httpRes, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdDelete(context.Background(), registrySourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrySourcesAPIService SourcesRegistryRegistrySourceIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var registrySourceId string

		resp, httpRes, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdGet(context.Background(), registrySourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrySourcesAPIService SourcesRegistryRegistrySourceIdPatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var registrySourceId string

		resp, httpRes, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdPatch(context.Background(), registrySourceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrySourcesAPIService WorkspacesWorkspaceIdSourcesRegistryGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspaceId string

		resp, httpRes, err := apiClient.RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryGet(context.Background(), workspaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RegistrySourcesAPIService WorkspacesWorkspaceIdSourcesRegistryPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspaceId string

		resp, httpRes, err := apiClient.RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryPost(context.Background(), workspaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
