# \RegistrySourcesAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourcesRegistryRegistrySourceIdComposeFileGet**](RegistrySourcesAPI.md#SourcesRegistryRegistrySourceIdComposeFileGet) | **Get** /sources/registry/{registry_source_id}/compose_file | 
[**SourcesRegistryRegistrySourceIdDelete**](RegistrySourcesAPI.md#SourcesRegistryRegistrySourceIdDelete) | **Delete** /sources/registry/{registry_source_id} | 
[**SourcesRegistryRegistrySourceIdGet**](RegistrySourcesAPI.md#SourcesRegistryRegistrySourceIdGet) | **Get** /sources/registry/{registry_source_id} | 
[**SourcesRegistryRegistrySourceIdPatch**](RegistrySourcesAPI.md#SourcesRegistryRegistrySourceIdPatch) | **Patch** /sources/registry/{registry_source_id} | 
[**WorkspacesWorkspaceIdSourcesRegistryGet**](RegistrySourcesAPI.md#WorkspacesWorkspaceIdSourcesRegistryGet) | **Get** /workspaces/{workspace_id}/sources/registry | 
[**WorkspacesWorkspaceIdSourcesRegistryPost**](RegistrySourcesAPI.md#WorkspacesWorkspaceIdSourcesRegistryPost) | **Post** /workspaces/{workspace_id}/sources/registry | 



## SourcesRegistryRegistrySourceIdComposeFileGet

> ComposeFile SourcesRegistryRegistrySourceIdComposeFileGet(ctx, registrySourceId).Execute()



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
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdComposeFileGet(context.Background(), registrySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdComposeFileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistryRegistrySourceIdComposeFileGet`: ComposeFile
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdComposeFileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistryRegistrySourceIdComposeFileGetRequest struct via the builder pattern


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


## SourcesRegistryRegistrySourceIdDelete

> RegistrySource SourcesRegistryRegistrySourceIdDelete(ctx, registrySourceId).Execute()



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
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdDelete(context.Background(), registrySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistryRegistrySourceIdDelete`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistryRegistrySourceIdDeleteRequest struct via the builder pattern


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


## SourcesRegistryRegistrySourceIdGet

> RegistrySource SourcesRegistryRegistrySourceIdGet(ctx, registrySourceId).Execute()



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
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdGet(context.Background(), registrySourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistryRegistrySourceIdGet`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistryRegistrySourceIdGetRequest struct via the builder pattern


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


## SourcesRegistryRegistrySourceIdPatch

> RegistrySource SourcesRegistryRegistrySourceIdPatch(ctx, registrySourceId).UpdateRegistrySource(updateRegistrySource).Execute()



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
  restart: always", []string{"ExtraRepositories_example"}, true, false) // UpdateRegistrySource | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegistrySourcesAPI.SourcesRegistryRegistrySourceIdPatch(context.Background(), registrySourceId).UpdateRegistrySource(updateRegistrySource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdPatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistryRegistrySourceIdPatch`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.SourcesRegistryRegistrySourceIdPatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistryRegistrySourceIdPatchRequest struct via the builder pattern


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


## WorkspacesWorkspaceIdSourcesRegistryGet

> RegistrySourceList WorkspacesWorkspaceIdSourcesRegistryGet(ctx, workspaceId).Execute()



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
	resp, r, err := apiClient.RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryGet(context.Background(), workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdSourcesRegistryGet`: RegistrySourceList
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdSourcesRegistryGetRequest struct via the builder pattern


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


## WorkspacesWorkspaceIdSourcesRegistryPost

> RegistrySource WorkspacesWorkspaceIdSourcesRegistryPost(ctx, workspaceId).CreateRegistrySource(createRegistrySource).Execute()



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
	resp, r, err := apiClient.RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryPost(context.Background(), workspaceId).CreateRegistrySource(createRegistrySource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspacesWorkspaceIdSourcesRegistryPost`: RegistrySource
	fmt.Fprintf(os.Stdout, "Response from `RegistrySourcesAPI.WorkspacesWorkspaceIdSourcesRegistryPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 

### Other Parameters

Other parameters are passed through a pointer to a apiWorkspacesWorkspaceIdSourcesRegistryPostRequest struct via the builder pattern


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

