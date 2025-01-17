/*
Miru API

Testing GenericErrorsAPIService

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

func Test_openapi_GenericErrorsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GenericErrorsAPIService ExampleErrors400Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GenericErrorsAPI.ExampleErrors400Get(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenericErrorsAPIService ExampleErrors401Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GenericErrorsAPI.ExampleErrors401Get(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenericErrorsAPIService ExampleErrors403Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GenericErrorsAPI.ExampleErrors403Get(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GenericErrorsAPIService ExampleErrors404Get", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		httpRes, err := apiClient.GenericErrorsAPI.ExampleErrors404Get(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
