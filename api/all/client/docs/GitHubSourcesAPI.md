# \GitHubSourcesAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SourcesGithubGithubSourceIdCommitsGet**](GitHubSourcesAPI.md#SourcesGithubGithubSourceIdCommitsGet) | **Get** /sources/github/{github_source_id}/commits | 



## SourcesGithubGithubSourceIdCommitsGet

> GitHubSourceCommits SourcesGithubGithubSourceIdCommitsGet(ctx, workspaceId, githubSourceId).Execute()



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
	githubSourceId := "githubSourceId_example" // string | The unique identifier of the github source

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GitHubSourcesAPI.SourcesGithubGithubSourceIdCommitsGet(context.Background(), workspaceId, githubSourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GitHubSourcesAPI.SourcesGithubGithubSourceIdCommitsGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SourcesGithubGithubSourceIdCommitsGet`: GitHubSourceCommits
	fmt.Fprintf(os.Stdout, "Response from `GitHubSourcesAPI.SourcesGithubGithubSourceIdCommitsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The unique identifier of the workspace | 
**githubSourceId** | **string** | The unique identifier of the github source | 

### Other Parameters

Other parameters are passed through a pointer to a apiSourcesGithubGithubSourceIdCommitsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GitHubSourceCommits**](GitHubSourceCommits.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

