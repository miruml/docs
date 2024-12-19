# \ArtifactsAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtifactsArtifactIdDeploymentFilesGet**](ArtifactsAPI.md#ArtifactsArtifactIdDeploymentFilesGet) | **Get** /artifacts/{artifact_id}/deployment_files | 
[**ArtifactsArtifactIdGet**](ArtifactsAPI.md#ArtifactsArtifactIdGet) | **Get** /artifacts/{artifact_id} | 
[**GroupsGroupIdArtifactsGet**](ArtifactsAPI.md#GroupsGroupIdArtifactsGet) | **Get** /groups/{group_id}/artifacts | 
[**SourcesGithubGithubSourceIdArtifactsPost**](ArtifactsAPI.md#SourcesGithubGithubSourceIdArtifactsPost) | **Post** /sources/github/{github_source_id}/artifacts | 
[**SourcesRegistryRegistrySourceIdArtifactsPost**](ArtifactsAPI.md#SourcesRegistryRegistrySourceIdArtifactsPost) | **Post** /sources/registry/{registry_source_id}/artifacts | 



## ArtifactsArtifactIdDeploymentFilesGet

> ArtifactDeploymentFiles ArtifactsArtifactIdDeploymentFilesGet(ctx, artifactId).Execute()



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
	artifactId := "artifactId_example" // string | The unique identifier of the artifact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.ArtifactsArtifactIdDeploymentFilesGet(context.Background(), artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsArtifactIdDeploymentFilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArtifactsArtifactIdDeploymentFilesGet`: ArtifactDeploymentFiles
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ArtifactsArtifactIdDeploymentFilesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactId** | **string** | The unique identifier of the artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsArtifactIdDeploymentFilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ArtifactDeploymentFiles**](ArtifactDeploymentFiles.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArtifactsArtifactIdGet

> Artifact ArtifactsArtifactIdGet(ctx, artifactId).Execute()



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
	artifactId := "artifactId_example" // string | The unique identifier of the artifact

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.ArtifactsArtifactIdGet(context.Background(), artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsArtifactIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArtifactsArtifactIdGet`: Artifact
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ArtifactsArtifactIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactId** | **string** | The unique identifier of the artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsArtifactIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Artifact**](Artifact.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GroupsGroupIdArtifactsGet

> GroupArtifactList GroupsGroupIdArtifactsGet(ctx, groupId).OnlyStaged(onlyStaged).Execute()



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
	groupId := "groupId_example" // string | The unique identifier of the group
	onlyStaged := true // bool | Whether to return only staged artifacts (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.GroupsGroupIdArtifactsGet(context.Background(), groupId).OnlyStaged(onlyStaged).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.GroupsGroupIdArtifactsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GroupsGroupIdArtifactsGet`: GroupArtifactList
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.GroupsGroupIdArtifactsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | The unique identifier of the group | 

### Other Parameters

Other parameters are passed through a pointer to a apiGroupsGroupIdArtifactsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **onlyStaged** | **bool** | Whether to return only staged artifacts | 

### Return type

[**GroupArtifactList**](GroupArtifactList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourcesGithubGithubSourceIdArtifactsPost

> SourcesGithubGithubSourceIdArtifactsPost(ctx, githubSourceId).CreateGitHubSourceArtifact(createGitHubSourceArtifact).Execute()



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
	githubSourceId := "githubSourceId_example" // string | The unique identifier of the github source
	createGitHubSourceArtifact := *openapiclient.NewCreateGitHubSourceArtifact(false, "94d8a2681f1ef81b0362cdb77bf37691bc1d2b03") // CreateGitHubSourceArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ArtifactsAPI.SourcesGithubGithubSourceIdArtifactsPost(context.Background(), githubSourceId).CreateGitHubSourceArtifact(createGitHubSourceArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.SourcesGithubGithubSourceIdArtifactsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**githubSourceId** | **string** | The unique identifier of the github source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesGithubGithubSourceIdArtifactsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createGitHubSourceArtifact** | [**CreateGitHubSourceArtifact**](CreateGitHubSourceArtifact.md) |  | 

### Return type

 (empty response body)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourcesRegistryRegistrySourceIdArtifactsPost

> Artifact SourcesRegistryRegistrySourceIdArtifactsPost(ctx, registrySourceId).CreateRegistrySourceArtifact(createRegistrySourceArtifact).Execute()



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
	createRegistrySourceArtifact := *openapiclient.NewCreateRegistrySourceArtifact(false, []openapiclient.CreateRegistrySourceArtifactImagesInner{*openapiclient.NewCreateRegistrySourceArtifactImagesInner("cont_repo_1234", "sha256:1234567890", "v1.3.4")}) // CreateRegistrySourceArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.SourcesRegistryRegistrySourceIdArtifactsPost(context.Background(), registrySourceId).CreateRegistrySourceArtifact(createRegistrySourceArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.SourcesRegistryRegistrySourceIdArtifactsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistryRegistrySourceIdArtifactsPost`: Artifact
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.SourcesRegistryRegistrySourceIdArtifactsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistryRegistrySourceIdArtifactsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRegistrySourceArtifact** | [**CreateRegistrySourceArtifact**](CreateRegistrySourceArtifact.md) |  | 

### Return type

[**Artifact**](Artifact.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

