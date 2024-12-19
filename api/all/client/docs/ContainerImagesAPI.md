# \ContainerImagesAPI

All URIs are relative to *http://localhost:8080/internal/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RepositoriesContainerContainerRepositoryIdImagesExternalGet**](ContainerImagesAPI.md#RepositoriesContainerContainerRepositoryIdImagesExternalGet) | **Get** /repositories/container/{container_repository_id}/images/external | 



## RepositoriesContainerContainerRepositoryIdImagesExternalGet

> ExternalContainerImageList RepositoriesContainerContainerRepositoryIdImagesExternalGet(ctx, containerRepositoryId).PageSize(pageSize).Execute()



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
	pageSize := int32(56) // int32 | The number of items to return per page (optional) (default to 25)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ContainerImagesAPI.RepositoriesContainerContainerRepositoryIdImagesExternalGet(context.Background(), containerRepositoryId).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ContainerImagesAPI.RepositoriesContainerContainerRepositoryIdImagesExternalGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RepositoriesContainerContainerRepositoryIdImagesExternalGet`: ExternalContainerImageList
	fmt.Fprintf(os.Stdout, "Response from `ContainerImagesAPI.RepositoriesContainerContainerRepositoryIdImagesExternalGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**containerRepositoryId** | **string** | The unique identifier of the container repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiRepositoriesContainerContainerRepositoryIdImagesExternalGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pageSize** | **int32** | The number of items to return per page | [default to 25]

### Return type

[**ExternalContainerImageList**](ExternalContainerImageList.md)

### Authorization

[ClerkAuth](../README.md#ClerkAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

