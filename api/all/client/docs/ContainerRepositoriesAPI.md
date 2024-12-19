# \ContainerRepositoriesAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet**](ContainerRepositoriesAPI.md#IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet) | **Get** /integrations/dockerhub/{dockerhub_integration_id}/repositories/external | 
[**IntegrationsGarGarIntegrationIdRepositoriesExternalGet**](ContainerRepositoriesAPI.md#IntegrationsGarGarIntegrationIdRepositoriesExternalGet) | **Get** /integrations/gar/{gar_integration_id}/repositories/external | 
[**RepositoriesContainerContainerRepositoryIdGet**](ContainerRepositoriesAPI.md#RepositoriesContainerContainerRepositoryIdGet) | **Get** /repositories/container/{container_repository_id} | 
[**WorkspacesWorkspaceIdRepositoriesContainerExternalGet**](ContainerRepositoriesAPI.md#WorkspacesWorkspaceIdRepositoriesContainerExternalGet) | **Get** /workspaces/{workspace_id}/repositories/container/external | 



## IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet

> ExternalContainerRepositoryList IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(ctx, dockerhubIntegrationId).PageSize(pageSize).Execute()



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
	pageSize := int32(56) // int32 | The number of items to return per page (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRepositoriesAPI.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet(context.Background(), dockerhubIntegrationId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRepositoriesAPI.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet`: ExternalContainerRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `ContainerRepositoriesAPI.IntegrationsDockerhubDockerhubIntegrationIdRepositoriesExternalGet`: %v\n", resp)
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

 **pageSize** | **int32** | The number of items to return per page | [default to 25]

### Return type

[**ExternalContainerRepositoryList**](ExternalContainerRepositoryList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IntegrationsGarGarIntegrationIdRepositoriesExternalGet

> ExternalContainerRepositoryList IntegrationsGarGarIntegrationIdRepositoriesExternalGet(ctx, garIntegrationId).PageSize(pageSize).Execute()



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
	pageSize := int32(56) // int32 | The number of items to return per page (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRepositoriesAPI.IntegrationsGarGarIntegrationIdRepositoriesExternalGet(context.Background(), garIntegrationId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRepositoriesAPI.IntegrationsGarGarIntegrationIdRepositoriesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IntegrationsGarGarIntegrationIdRepositoriesExternalGet`: ExternalContainerRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `ContainerRepositoriesAPI.IntegrationsGarGarIntegrationIdRepositoriesExternalGet`: %v\n", resp)
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

 **pageSize** | **int32** | The number of items to return per page | [default to 25]

### Return type

[**ExternalContainerRepositoryList**](ExternalContainerRepositoryList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepositoriesContainerContainerRepositoryIdGet

> ContainerRepository RepositoriesContainerContainerRepositoryIdGet(ctx, containerRepositoryId).Execute()



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
	containerRepositoryId := "containerRepositoryId_example" // string | The unique identifier of the container repository

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRepositoriesAPI.RepositoriesContainerContainerRepositoryIdGet(context.Background(), containerRepositoryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRepositoriesAPI.RepositoriesContainerContainerRepositoryIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerRepositoryIdGet`: ContainerRepository
	fmt.Fprintf(os.Stdout, "Response from `ContainerRepositoriesAPI.RepositoriesContainerContainerRepositoryIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerRepositoryId** | **string** | The unique identifier of the container repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerRepositoryIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerRepository**](ContainerRepository.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdRepositoriesContainerExternalGet

> ExternalContainerRepositoryList WorkspacesWorkspaceIdRepositoriesContainerExternalGet(ctx, workspaceId).PageSize(pageSize).Execute()



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
	pageSize := int32(56) // int32 | The number of items to return per page (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerRepositoriesAPI.WorkspacesWorkspaceIdRepositoriesContainerExternalGet(context.Background(), workspaceId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerRepositoriesAPI.WorkspacesWorkspaceIdRepositoriesContainerExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdRepositoriesContainerExternalGet`: ExternalContainerRepositoryList
	fmt.Fprintf(os.Stdout, "Response from `ContainerRepositoriesAPI.WorkspacesWorkspaceIdRepositoriesContainerExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdRepositoriesContainerExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to return per page | [default to 25]

### Return type

[**ExternalContainerRepositoryList**](ExternalContainerRepositoryList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

