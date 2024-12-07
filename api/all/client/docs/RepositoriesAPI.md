# \RepositoriesAPI

All URIs are relative to *https://api.miruml.com/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet**](RepositoriesAPI.md#IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet) | **Get** /integrations/dockerhub/{dockerhub_integration_id}/repositories/external | 
[**IntegrationsGarGarIntegrationIdRepositoriesExternalGet**](RepositoriesAPI.md#IntegrationsGarGarIntegrationIdRepositoriesExternalGet) | **Get** /integrations/gar/{gar_integration_id}/repositories/external | 
[**RepositoriesRepositoryIdGet**](RepositoriesAPI.md#RepositoriesRepositoryIdGet) | **Get** /repositories/{repository_id} | 
[**RepositoriesRepositoryIdImagesExternalGet**](RepositoriesAPI.md#RepositoriesRepositoryIdImagesExternalGet) | **Get** /repositories/{repository_id}/images/external | 
[**WorkspacesWorkspaceIdRepositoriesExternalGet**](RepositoriesAPI.md#WorkspacesWorkspaceIdRepositoriesExternalGet) | **Get** /workspaces/{workspace_id}/repositories/external | 



## IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet

> ExternalRepositoryList IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(ctx, dockerhubIntegrationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	dockerhubIntegrationId := "dockerhubIntegrationId_example" // string | The unique identifier of the dockerhub integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(context.Background(), dockerhubIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet`: ExternalRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dockerhubIntegrationId** | **string** | The unique identifier of the dockerhub integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalRepositoryList**](ExternalRepositoryList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGarGarIntegrationIdRepositoriesExternalGet

> ExternalRepositoryList IntegrationsGarGarIntegrationIdRepositoriesExternalGet(ctx, garIntegrationId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	garIntegrationId := "garIntegrationId_example" // string | The unique identifier of the google artifact registry integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.IntegrationsGarGarIntegrationIdRepositoriesExternalGet(context.Background(), garIntegrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.IntegrationsGarGarIntegrationIdRepositoriesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsGarGarIntegrationIdRepositoriesExternalGet`: ExternalRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.IntegrationsGarGarIntegrationIdRepositoriesExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**garIntegrationId** | **string** | The unique identifier of the google artifact registry integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiIntegrationsGarGarIntegrationIdRepositoriesExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalRepositoryList**](ExternalRepositoryList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesRepositoryIdGet

> Repository RepositoriesRepositoryIdGet(ctx, repositoryId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	repositoryId := "repositoryId_example" // string | The unique identifier of the repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.RepositoriesRepositoryIdGet(context.Background(), repositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.RepositoriesRepositoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRepositoryIdGet`: Repository
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.RepositoriesRepositoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **string** | The unique identifier of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRepositoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Repository**](Repository.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesRepositoryIdImagesExternalGet

> ExternalImageList RepositoriesRepositoryIdImagesExternalGet(ctx, repositoryId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	repositoryId := "repositoryId_example" // string | The unique identifier of the repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.RepositoriesRepositoryIdImagesExternalGet(context.Background(), repositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.RepositoriesRepositoryIdImagesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesRepositoryIdImagesExternalGet`: ExternalImageList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.RepositoriesRepositoryIdImagesExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **string** | The unique identifier of the repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesRepositoryIdImagesExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalImageList**](ExternalImageList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdRepositoriesExternalGet

> ExternalRepositoryList WorkspacesWorkspaceIdRepositoriesExternalGet(ctx, workspaceId).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	workspaceId := "workspaceId_example" // string | The unique identifier of the workspace

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RepositoriesAPI.WorkspacesWorkspaceIdRepositoriesExternalGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RepositoriesAPI.WorkspacesWorkspaceIdRepositoriesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdRepositoriesExternalGet`: ExternalRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `RepositoriesAPI.WorkspacesWorkspaceIdRepositoriesExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdRepositoriesExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExternalRepositoryList**](ExternalRepositoryList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

