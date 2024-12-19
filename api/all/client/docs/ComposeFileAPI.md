# \ComposeFileAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WorkspaceWorkspaceIdComposeVerifyPost**](ComposeFileAPI.md#WorkspaceWorkspaceIdComposeVerifyPost) | **Post** /workspace/{workspace_id}/compose/verify | 



## WorkspaceWorkspaceIdComposeVerifyPost

> VerifiedComposeFileResponse WorkspaceWorkspaceIdComposeVerifyPost(ctx).VerifyComposeFile(verifyComposeFile).Execute()



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
	verifyComposeFile := *openapiclient.NewVerifyComposeFile("compose_file", "name: miru
services:
  redis:
  image: redis:latest
  restart: always") // VerifyComposeFile | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ComposeFileAPI.WorkspaceWorkspaceIdComposeVerifyPost(context.Background()).VerifyComposeFile(verifyComposeFile).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ComposeFileAPI.WorkspaceWorkspaceIdComposeVerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkspaceWorkspaceIdComposeVerifyPost`: VerifiedComposeFileResponse
	fmt.Fprintf(os.Stdout, "Response from `ComposeFileAPI.WorkspaceWorkspaceIdComposeVerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkspaceWorkspaceIdComposeVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **verifyComposeFile** | [**VerifyComposeFile**](VerifyComposeFile.md) |  | 

### Return type

[**VerifiedComposeFileResponse**](VerifiedComposeFileResponse.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

