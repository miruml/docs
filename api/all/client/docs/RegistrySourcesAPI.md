# \RegistrySourcesAPI

All URIs are relative to *https://api.miruml.com/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourcesRegistriesRegistrySourceIdComposeFileGet**](RegistrySourcesAPI.md#SourcesRegistriesRegistrySourceIdComposeFileGet) | **Get** /sources/registries/{registry_source_id}/compose_file | 
[**SourcesRegistriesRegistrySourceIdGet**](RegistrySourcesAPI.md#SourcesRegistriesRegistrySourceIdGet) | **Get** /sources/registries/{registry_source_id} | 
[**SourcesRegistriesRegistrySourceIdPatch**](RegistrySourcesAPI.md#SourcesRegistriesRegistrySourceIdPatch) | **Patch** /sources/registries/{registry_source_id} | 
[**WorkspacesWorkspaceIdSourcesRegistriesGet**](RegistrySourcesAPI.md#WorkspacesWorkspaceIdSourcesRegistriesGet) | **Get** /workspaces/{workspace_id}/sources/registries | 
[**WorkspacesWorkspaceIdSourcesRegistriesPost**](RegistrySourcesAPI.md#WorkspacesWorkspaceIdSourcesRegistriesPost) | **Post** /workspaces/{workspace_id}/sources/registries | 



## SourcesRegistriesRegistrySourceIdComposeFileGet

> ComposeFile SourcesRegistriesRegistrySourceIdComposeFileGet(ctx, registrySourceId).Execute()



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
	registrySourceId := "registrySourceId_example" // string | The unique identifier of the registry source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdComposeFileGet(context.Background(), registrySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdComposeFileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistriesRegistrySourceIdComposeFileGet`: ComposeFile
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdComposeFileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistriesRegistrySourceIdComposeFileGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComposeFile**](ComposeFile.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourcesRegistriesRegistrySourceIdGet

> RegistrySource SourcesRegistriesRegistrySourceIdGet(ctx, registrySourceId).Execute()



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
	registrySourceId := "registrySourceId_example" // string | The unique identifier of the registry source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdGet(context.Background(), registrySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistriesRegistrySourceIdGet`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistriesRegistrySourceIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistrySource**](RegistrySource.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourcesRegistriesRegistrySourceIdPatch

> RegistrySource SourcesRegistriesRegistrySourceIdPatch(ctx, registrySourceId).UpdateRegistrySource(updateRegistrySource).Execute()



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
	registrySourceId := "registrySourceId_example" // string | The unique identifier of the registry source
	updateRegistrySource := *openapiclient.NewUpdateRegistrySource("Ad Display V2.2.1", "name: miru
services:
  redis:
  image: redis:latest
  restart: always", *openapiclient.NewUpdateRegistrySourceExtraRepositories(), true, false) // UpdateRegistrySource | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdPatch(context.Background(), registrySourceId).UpdateRegistrySource(updateRegistrySource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistriesRegistrySourceIdPatch`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistriesRegistrySourceIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistriesRegistrySourceIdPatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRegistrySource** | [**UpdateRegistrySource**](UpdateRegistrySource.md) |  | 

### Return type

[**RegistrySource**](RegistrySource.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdSourcesRegistriesGet

> RegistrySourceList WorkspacesWorkspaceIdSourcesRegistriesGet(ctx, workspaceId).Execute()



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
	resp, r, err := apiClient.RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistriesGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistriesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdSourcesRegistriesGet`: RegistrySourceList
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistriesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdSourcesRegistriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RegistrySourceList**](RegistrySourceList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkspacesWorkspaceIdSourcesRegistriesPost

> RegistrySource WorkspacesWorkspaceIdSourcesRegistriesPost(ctx, workspaceId).CreateRegistrySource(createRegistrySource).Execute()



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
	createRegistrySource := *openapiclient.NewCreateRegistrySource("Ad Display V2.2.1", "name: miru
services:
  redis:
  image: redis:latest
  restart: always", []string{"ExtraRepositories_example"}, true, false) // CreateRegistrySource | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistriesPost(context.Background(), workspaceId).CreateRegistrySource(createRegistrySource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistriesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdSourcesRegistriesPost`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistriesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdSourcesRegistriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRegistrySource** | [**CreateRegistrySource**](CreateRegistrySource.md) |  | 

### Return type

[**RegistrySource**](RegistrySource.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

