# \ArtifactsAPI

All URIs are relative to *https://api.miruml.com/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArtifactsArtifactIdComposeFileGet**](ArtifactsAPI.md#ArtifactsArtifactIdComposeFileGet) | **Get** /artifacts/{artifact_id}/compose_file | 
[**ArtifactsArtifactIdGet**](ArtifactsAPI.md#ArtifactsArtifactIdGet) | **Get** /artifacts/{artifact_id} | 
[**SourcesGithubGithubSourceIdArtifactsPost**](ArtifactsAPI.md#SourcesGithubGithubSourceIdArtifactsPost) | **Post** /sources/github/{github_source_id}/artifacts | 
[**SourcesRegistriesRegistrySourceIdArtifactsPost**](ArtifactsAPI.md#SourcesRegistriesRegistrySourceIdArtifactsPost) | **Post** /sources/registries/{registry_source_id}/artifacts | 



## ArtifactsArtifactIdComposeFileGet

> ComposeFile ArtifactsArtifactIdComposeFileGet(ctx, artifactId).Execute()



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
	resp, r, err := apiClient.ArtifactsAPI.ArtifactsArtifactIdComposeFileGet(context.Background(), artifactId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.ArtifactsArtifactIdComposeFileGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArtifactsArtifactIdComposeFileGet`: ComposeFile
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.ArtifactsArtifactIdComposeFileGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**artifactId** | **string** | The unique identifier of the artifact | 

### Other Parameters

Other parameters are passed through a pointer to a apiArtifactsArtifactIdComposeFileGetRequest struct via the builder pattern


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


## SourcesGithubGithubSourceIdArtifactsPost

> Artifact SourcesGithubGithubSourceIdArtifactsPost(ctx, githubSourceId).CreateGitHubSourceArtifact(createGitHubSourceArtifact).Execute()



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
	createGitHubSourceArtifact := *openapiclient.NewCreateGitHubSourceArtifact("94d8a2681f1ef81b0362cdb77bf37691bc1d2b03") // CreateGitHubSourceArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.SourcesGithubGithubSourceIdArtifactsPost(context.Background(), githubSourceId).CreateGitHubSourceArtifact(createGitHubSourceArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.SourcesGithubGithubSourceIdArtifactsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesGithubGithubSourceIdArtifactsPost`: Artifact
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.SourcesGithubGithubSourceIdArtifactsPost`: %v\n", resp)
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

[**Artifact**](Artifact.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SourcesRegistriesRegistrySourceIdArtifactsPost

> Artifact SourcesRegistriesRegistrySourceIdArtifactsPost(ctx, registrySourceId).CreateRegistrySourceArtifact(createRegistrySourceArtifact).Execute()



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
	createRegistrySourceArtifact := *openapiclient.NewCreateRegistrySourceArtifact([]openapiclient.CreateRegistrySourceArtifactImagesInner{*openapiclient.NewCreateRegistrySourceArtifactImagesInner("repo_1234", "v1.3.4")}) // CreateRegistrySourceArtifact | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ArtifactsAPI.SourcesRegistriesRegistrySourceIdArtifactsPost(context.Background(), registrySourceId).CreateRegistrySourceArtifact(createRegistrySourceArtifact).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ArtifactsAPI.SourcesRegistriesRegistrySourceIdArtifactsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesRegistriesRegistrySourceIdArtifactsPost`: Artifact
	fmt.Fprintf(os.Stdout, "Response from `ArtifactsAPI.SourcesRegistriesRegistrySourceIdArtifactsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrySourceId** | **string** | The unique identifier of the registry source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesRegistriesRegistrySourceIdArtifactsPostRequest struct via the builder pattern


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

